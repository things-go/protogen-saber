package main

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

func execute(g *protogen.GeneratedFile, s *serviceDesc) error {
	schedulerMethods := make([]*methodDesc, 0, len(s.Methods))
	// pattern constants
	for _, m := range s.Methods {
		if m.Pattern == "" {
			return fmt.Errorf("service %s(%s) pattern should be not empty", s.ServiceType, m.Name)
		}
		g.P("const ", patternConstant(s.ServiceType, m.Name), ` = "`, m.Pattern, `"`)
		if m.CronSpec != "" {
			g.P("const ", cronSpecConstant(s.ServiceType, m.Name), ` = "`, m.CronSpec, `"`)
			schedulerMethods = append(schedulerMethods, m)
		}
	}
	g.P()

	// server impl
	{
		// server interface
		g.P("type ", serverInterfaceName(s.ServiceType), " interface {")
		for _, m := range s.Methods {
			g.P(m.Comment)
			g.P(serverMethodDefinition(g, m))
		}
		g.P("}")
		g.P()
		// server factory
		g.P("func Register", s.ServiceType, "TaskHandler(mux *", g.QualifiedGoIdent(asynqPackage.Ident("ServeMux")), ", srv ", serverInterfaceName(s.ServiceType), ", opts ...", g.QualifiedGoIdent(asynqAuxiliaryPackage.Ident("HandlerOption")), ") {")
		g.P("settings :=", g.QualifiedGoIdent(asynqAuxiliaryPackage.Ident("NewHandlerSettings()")))
		g.P("for _, opt := range opts {")
		g.P("opt(settings)")
		g.P("}")
		for _, m := range s.Methods {
			g.P("mux.HandleFunc(", patternConstant(s.ServiceType, m.Name), ", ", serviceHandlerMethodName(s.ServiceType, m.Name), "(srv, settings))")
		}
		g.P("}")
		g.P()

		// server handler
		for _, m := range s.Methods {
			g.P("func ", serviceHandlerMethodName(s.ServiceType, m.Name), "(srv ", serverInterfaceName(s.ServiceType), ", settings *", g.QualifiedGoIdent(asynqAuxiliaryPackage.Ident("HandlerSettings")), ") ", asynqHandler(g, true), " {")
			{ // closure
				g.P("return ", asynqHandler(g, false), " {")
				g.P("var in ", m.Request)
				g.P()
				g.P("if err := settings.UnmarshalBinary(task.Payload(), &in); err != nil {")
				g.P("return err")
				g.P("}")
				g.P("return srv.", m.Name, "(ctx, &in)")
				g.P("}")
			}
			g.P("}")
			g.P()
		}
	}
	g.P()
	//  client impl
	{
		// client interface
		g.P("type ", clientInterfaceName(s.ServiceType), " interface {")
		for _, m := range s.Methods {
			g.P(m.Comment)
			g.P(clientMethodDefinition(g, m, true))
		}
		g.P("}")
		g.P()

		// client impl
		g.P("type ", clientImplStructName(s.ServiceType), " struct {")
		g.P("cc *", g.QualifiedGoIdent(asynqPackage.Ident("Client")))
		g.P("settings *", g.QualifiedGoIdent(asynqAuxiliaryPackage.Ident("ClientSettings")))
		g.P("}")
		g.P()
		// client factory
		g.P("// ", clientFactoryMethodName(s.ServiceType), " new client. use default proto.Marhsal.")
		g.P("func ", clientFactoryMethodName(s.ServiceType), " (client *", g.QualifiedGoIdent(asynqPackage.Ident("Client")), ", opts ...", g.QualifiedGoIdent(asynqAuxiliaryPackage.Ident("ClientOption")), ") ", clientInterfaceName(s.ServiceType), " {")
		{ // closure
			g.P("settings := ", g.QualifiedGoIdent(asynqAuxiliaryPackage.Ident("NewClientSettings()")))
			g.P("for _, opt := range opts {")
			g.P("opt(settings)")
			g.P("}")
			g.P("return &", clientImplStructName(s.ServiceType), " {")
			g.P("cc: client,")
			g.P("settings: settings,")
			g.P("}")
		}
		g.P("}")
		g.P()
		// client method
		for _, m := range s.Methods {
			g.P(m.Comment)
			g.P("func (c *", clientImplStructName(s.ServiceType), ")", clientMethodDefinition(g, m, false), " {")
			g.P("payload, err := c.settings.MarshalBinary(in)")
			g.P("if err != nil {")
			g.P("return nil, err")
			g.P("}")
			g.P("task := ", g.QualifiedGoIdent(asynqPackage.Ident("NewTask")), "(", patternConstant(s.ServiceType, m.Name), ", payload, opts...)")
			g.P("return c.cc.Enqueue(task)")
			g.P("}")
			g.P()
		}
	}
	if len(schedulerMethods) > 0 {
		for _, m := range schedulerMethods {
			// _, err = scheduler.Register(
			// 	mq_bank.CronSpec_Bank_BankTransferQueryScheduler,
			// 	asynq.NewTask(mq_bank.Pattern_Bank_BankTransferQueryScheduler, emptyData),
			// )
			g.P("func RegisterScheduler", m.Name,
				"(scheduler *", g.QualifiedGoIdent(asynqPackage.Ident("Scheduler")),
				", in *", m.Request,
				", settings *", g.QualifiedGoIdent(asynqAuxiliaryPackage.Ident("ClientSettings")),
				", opts ..."+g.QualifiedGoIdent(asynqPackage.Ident("Option")), ") (entryId string, err error) {")
			g.P("var payload []byte")
			g.P()
			g.P("if settings.MarshalBinary != nil {")
			g.P("payload, err = settings.MarshalBinary(in)")
			g.P("} else {")
			g.P("payload, err = ", g.QualifiedGoIdent(protoPackage.Ident("Marshal")), "(in)")
			g.P("}")
			g.P("if err != nil {")
			g.P("return \"\", err")
			g.P("}")
			g.P("return scheduler.Register(", cronSpecConstant(s.ServiceType, m.Name), ", ", g.QualifiedGoIdent(asynqPackage.Ident("NewTask")), "(", patternConstant(s.ServiceType, m.Name), ", payload, opts...))")
			g.P("}")
			g.P()
		}
	}

	return nil
}

func patternConstant(serviceType, name string) string {
	return "Pattern_" + serviceType + "_" + name
}

func cronSpecConstant(serviceType, name string) string {
	return "CronSpec_" + serviceType + "_" + name
}

func serverInterfaceName(serviceType string) string {
	return serviceType + "TaskHandler"
}
func serverMethodDefinition(g *protogen.GeneratedFile, m *methodDesc) string {
	return m.Name + "(" + g.QualifiedGoIdent(contextPackage.Ident("Context")) + ", *" + m.Request + ") error"
}
func serviceHandlerMethodName(serviceType, name string) string {
	return "_" + serviceType + "_" + name + "_Task_Handler"
}

func clientInterfaceName(serviceType string) string {
	return serviceType + "TaskClient"
}
func clientMethodDefinition(g *protogen.GeneratedFile, m *methodDesc, isDeclaration bool) string {
	ctxParam := ""
	inParam := ""
	optsParam := ""
	if !isDeclaration {
		ctxParam = "ctx"
		inParam = "in"
		optsParam = "opts"
	}
	return m.Name + "(" + ctxParam + " " +
		g.QualifiedGoIdent(contextPackage.Ident("Context")) +
		", " + inParam + " *" + m.Request +
		", " + optsParam + " ..." + g.QualifiedGoIdent(asynqPackage.Ident("Option")) +
		") (*" + g.QualifiedGoIdent(asynqPackage.Ident("TaskInfo")) + ", error)"
}
func clientImplStructName(serviceType string) string {
	return serviceType + "TaskClientImpl"
}
func clientFactoryMethodName(serviceType string) string {
	return "New" + serviceType + "TaskClient"
}

func asynqHandler(g *protogen.GeneratedFile, isDeclaration bool) string {
	ctxParam := ""
	taskParam := ""
	if !isDeclaration {
		ctxParam = "ctx"
		taskParam = "task"
	}
	return "func(" + ctxParam + " " + g.QualifiedGoIdent(contextPackage.Ident("Context")) + ", " +
		taskParam + " *" + g.QualifiedGoIdent(asynqPackage.Ident("Task")) + ") error"
}
