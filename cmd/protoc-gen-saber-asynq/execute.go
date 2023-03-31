package main

import (
	"google.golang.org/protobuf/compiler/protogen"
)

type serviceDesc struct {
	ServiceType string // Greeter
	ServiceName string // helloworld.Greeter
	Metadata    string // api/v1/helloworld.proto
	Methods     []*methodDesc
}

type methodDesc struct {
	Name    string // 方法名
	Num     int    // 方法号, not used
	Request string // 请求结构
	Reply   string // 回复结构
	Comment string // 方法注释
	// asynq rule
	Pattern string // 匹配器
}

func execute(g *protogen.GeneratedFile, s *serviceDesc) error {
	// pattern constants
	for _, m := range s.Methods {
		g.P("const ", patternConstant(s.ServiceType, m.Name), ` = "`, m.Pattern, `"`)
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
		g.P("// UnmarshalBinary parses the binary data and stores the result")
		g.P("// in the value pointed to by v.")
		g.P("UnmarshalBinary([]byte, any) error")
		g.P("}")
		g.P()
		// server unimplemented with default unmarshaler.
		g.P("type Unimplemented", s.ServiceType, "TaskHandlerImpl struct {}")
		g.P()
		g.P("func (*Unimplemented", s.ServiceType, "TaskHandlerImpl) UnmarshalBinary(b []byte, v any) error {")
		g.P("return ", g.QualifiedGoIdent(protoPackage.Ident("Unmarshal")), "(b, v.(", g.QualifiedGoIdent(protoPackage.Ident("Message")), "))")
		g.P("}")
		g.P()
		// server factory
		g.P("func Register", s.ServiceType, "TaskHandler(mux *", g.QualifiedGoIdent(asynqPackage.Ident("ServeMux")), ", srv ", serverInterfaceName(s.ServiceType), ") {")
		for _, m := range s.Methods {
			g.P("mux.HandleFunc(", patternConstant(s.ServiceType, m.Name), ", ", serviceHandlerMethodName(s.ServiceType, m.Name), "(srv))")
		}
		g.P("}")
		g.P()

		// server handler
		for _, m := range s.Methods {
			g.P("func ", serviceHandlerMethodName(s.ServiceType, m.Name), "(srv ", serverInterfaceName(s.ServiceType), ") ", asynqHandler(g, true), " {")
			{ // closure
				g.P("return ", asynqHandler(g, false), " {")
				g.P("var in ", m.Request)
				g.P()
				g.P("if err := srv.UnmarshalBinary(task.Payload(), &in); err != nil {")
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
		g.P("// SetMarshaler set marshal the binary encoding of v function.")
		g.P("SetMarshaler(func(any) ([]byte, error)) ", clientInterfaceName(s.ServiceType))
		for _, m := range s.Methods {
			g.P(m.Comment)
			g.P(clientMethodDefinition(g, m, true))
		}
		g.P("}")
		g.P()

		// client impl
		g.P("type ", clientImplStructName(s.ServiceType), " struct {")
		g.P("cc *", g.QualifiedGoIdent(asynqPackage.Ident("Client")))
		g.P("marshaler func(any) ([]byte, error)")
		g.P("}")
		g.P()
		// client factory
		g.P("// ", clientFactoryMethodName(s.ServiceType), " new client. use default proto.Marhsal.")
		g.P("func ", clientFactoryMethodName(s.ServiceType), " (client *", g.QualifiedGoIdent(asynqPackage.Ident("Client")), ") ", clientInterfaceName(s.ServiceType), " {")
		{ // closure
			g.P("return &", clientImplStructName(s.ServiceType), " {")
			g.P("cc: client,")
			g.P("marshaler: func(v any) ([]byte, error) {")
			g.P("return ", g.QualifiedGoIdent(protoPackage.Ident("Marshal")), "(v.(", g.QualifiedGoIdent(protoPackage.Ident("Message")), "))")
			g.P("},")
			g.P("}")
		}
		g.P("}")
		g.P()
		// client method
		g.P("// SetMarshaler set marshal the binary encoding of v function.")
		g.P("func (c *", clientImplStructName(s.ServiceType), ") SetMarshaler(marshaler func(any) ([]byte, error)) ", clientInterfaceName(s.ServiceType), " {")
		g.P("if marshaler != nil {")
		g.P("c.marshaler = marshaler")
		g.P("}")
		g.P("return c")
		g.P("}")
		g.P()
		for _, m := range s.Methods {
			g.P(m.Comment)
			g.P("func (c *", clientImplStructName(s.ServiceType), ")", clientMethodDefinition(g, m, false), " {")
			g.P("payload, err := c.marshaler(in)")
			g.P("if err != nil {")
			g.P("return nil, err")
			g.P("}")
			g.P("task := ", g.QualifiedGoIdent(asynqPackage.Ident("NewTask")), "(", patternConstant(s.ServiceType, m.Name), ", payload, opts...)")
			g.P("taskInfo, err := c.cc.Enqueue(task)")
			g.P("if err != nil {")
			g.P("return nil, err")
			g.P("}")
			g.P("return taskInfo, nil")
			g.P("}")
			g.P()
		}
	}

	return nil
}

func patternConstant(serviceType, name string) string {
	return "Pattern_" + serviceType + "_" + name
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
