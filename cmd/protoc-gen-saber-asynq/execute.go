package main

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

func execute(g *protogen.GeneratedFile, s *serviceDesc) error {
	schedulerMethods := make([]*methodDesc, 0, len(s.Methods))
	// pattern constants
	for _, m := range s.Methods {
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
		if !args.DisableSaber {
			g.P("func Register", s.ServiceType, "TaskHandler(mux *", g.QualifiedGoIdent(asynqPackage.Ident("ServeMux")), ", srv ", serverInterfaceName(s.ServiceType), ", opts ...", g.QualifiedGoIdent(asynqAuxiliaryPackage.Ident("HandlerOption")), ") {")
			g.P("settings :=", g.QualifiedGoIdent(asynqAuxiliaryPackage.Ident("NewHandlerSettings(opts...)")))
		} else {
			g.P("func Register", s.ServiceType, "TaskHandler(mux *", g.QualifiedGoIdent(asynqPackage.Ident("ServeMux")), ", srv ", serverInterfaceName(s.ServiceType), ") {")
		}
		for _, m := range s.Methods {
			if !args.DisableSaber {
				g.P("mux.HandleFunc(", patternConstant(s.ServiceType, m.Name), ", ", serviceHandlerMethodName(s.ServiceType, m.Name), "(srv, settings))")
			} else {
				g.P("mux.HandleFunc(", patternConstant(s.ServiceType, m.Name), ", ", serviceHandlerMethodName(s.ServiceType, m.Name), "(srv))")
			}
		}
		g.P("}")
		g.P()

		// server handler
		for _, m := range s.Methods {
			if !args.DisableSaber {
				g.P("func ", serviceHandlerMethodName(s.ServiceType, m.Name), "(srv ", serverInterfaceName(s.ServiceType), ", settings *", g.QualifiedGoIdent(asynqAuxiliaryPackage.Ident("HandlerSettings")), ") ", asynqHandler(g, true), " {")
			} else {
				g.P("func ", serviceHandlerMethodName(s.ServiceType, m.Name), "(srv ", serverInterfaceName(s.ServiceType), ") ", asynqHandler(g, true), " {")
			}
			{ // closure
				g.P("return ", asynqHandler(g, false), " {")
				g.P("var in ", m.Request)
				g.P()
				if !args.DisableSaber {
					g.P("if err := settings.Unmarshaler.UnmarshalBinary(task.Payload(), &in); err != nil {")
				} else {
					g.P("if err := ", g.QualifiedGoIdent(codecPackage().Ident("Unmarshal")), "(task.Payload(), &in); err != nil {")
				}
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
		if !args.DisableSaber {
			g.P("settings *", g.QualifiedGoIdent(asynqAuxiliaryPackage.Ident("ClientSettings")))
		}
		g.P("}")
		g.P()
		// client factory
		g.P("// ", clientFactoryMethodName(s.ServiceType), " new client.")
		if !args.DisableSaber {
			g.P("func ", clientFactoryMethodName(s.ServiceType), " (client *", g.QualifiedGoIdent(asynqPackage.Ident("Client")), ", opts ...", g.QualifiedGoIdent(asynqAuxiliaryPackage.Ident("ClientOption")), ") ", clientInterfaceName(s.ServiceType), " {")
		} else {
			g.P("func ", clientFactoryMethodName(s.ServiceType), " (client *", g.QualifiedGoIdent(asynqPackage.Ident("Client")), ") ", clientInterfaceName(s.ServiceType), " {")
		}
		{ // closure
			if !args.DisableSaber {
				g.P("settings := ", g.QualifiedGoIdent(asynqAuxiliaryPackage.Ident("NewClientSettings(opts...)")))
			}
			g.P("return &", clientImplStructName(s.ServiceType), " {")
			g.P("cc: client,")
			if !args.DisableSaber {
				g.P("settings: settings,")
			}
			g.P("}")
		}
		g.P("}")
		g.P()
		// client method
		for _, m := range s.Methods {
			g.P(m.Comment)
			g.P("func (c *", clientImplStructName(s.ServiceType), ")", clientMethodDefinition(g, m, false), " {")
			if !args.DisableSaber {
				g.P("payload, err := c.settings.Marshaler.MarshalBinary(in)")
			} else {
				g.P("payload, err := ", g.QualifiedGoIdent(codecPackage().Ident("Marshal")), "(in)")
			}
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
			if !args.DisableSaber {
				g.P("func RegisterScheduler_", s.ServiceType, "_", m.Name,
					"(scheduler *", g.QualifiedGoIdent(asynqPackage.Ident("Scheduler")),
					", settings *", g.QualifiedGoIdent(asynqAuxiliaryPackage.Ident("ClientSettings")),
					", in *", m.Request,
					", opts ..."+g.QualifiedGoIdent(asynqPackage.Ident("Option")), ") (entryId string, err error) {")
				g.P("if settings == nil {")
				g.P("settings = ", g.QualifiedGoIdent(asynqAuxiliaryPackage.Ident("NewClientSettings()")))
				g.P("}")
				g.P("payload, err := settings.Marshaler.MarshalBinary(in)")
			} else {
				g.P("func RegisterScheduler_", s.ServiceType, "_", m.Name,
					"(scheduler *", g.QualifiedGoIdent(asynqPackage.Ident("Scheduler")),
					", in *", m.Request,
					", opts ..."+g.QualifiedGoIdent(asynqPackage.Ident("Option")), ") (entryId string, err error) {")
				g.P("payload, err := ", g.QualifiedGoIdent(codecPackage().Ident("Marshal")), "(in)")
			}
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

var codec = map[string]protogen.GoImportPath{
	"proto": protogen.GoImportPath("google.golang.org/protobuf/proto"),
	"json":  protogen.GoImportPath("encoding/json"),
}

func checkSupportedCodec() error {
	if args.Codec == "custom" && args.CodecPackage != "" {
		return nil
	}
	if _, ok := codec[args.Codec]; ok {
		return nil
	}
	keys := make([]string, 0, len(codec)+1)
	for k := range codec {
		keys = append(keys, k)
	}
	keys = append(keys, "custom")
	return fmt.Errorf("`codec` only supported: %v", keys)
}

func codecPackage() protogen.GoImportPath {
	if args.Codec == "custom" {
		return protogen.GoImportPath(args.CodecPackage)
	}
	if v, ok := codec[args.Codec]; ok {
		return v
	}
	return codec["proto"]
}
