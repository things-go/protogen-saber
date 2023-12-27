package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/things-go/protogen-saber/internal/protoutil"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

const deprecationComment = "// Deprecated: Do not use."

var (
	errorsPackage         = protogen.GoImportPath("errors")
	contextPackage        = protogen.GoImportPath("context")
	asynqPackage          = protogen.GoImportPath("github.com/hibiken/asynq")
	emptyPackage          = protogen.GoImportPath("google.golang.org/protobuf/types/known/emptypb")
	asynqAuxiliaryPackage = protogen.GoImportPath("github.com/things-go/protogen-saber/core/asynq_auxiliary")
	// protoPackage          = protogen.GoImportPath("google.golang.org/protobuf/proto")
	// jsonPackage           = protogen.GoImportPath("encoding/json")
)

var methodSets = make(map[string]int)

func runProtoGen(gen *protogen.Plugin) error {
	gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}
		generateFile(gen, f)
	}
	return nil
}

// generateFile generates a .gin.pb.go file.
func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Services) == 0 || (!hasHTTPRule(file.Services)) {
		return nil
	}
	filename := file.GeneratedFilenamePrefix + ".asynq.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-saber-asynq. DO NOT EDIT.")
	g.P("// versions:")
	g.P("//   - protoc-gen-saber-asynq ", version)
	g.P("//   - protoc                 ", protoutil.ProtocVersion(gen))
	if file.Proto.GetOptions().GetDeprecated() {
		g.P("// ", file.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// source: ", file.Desc.Path())
	}
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	generateFileContent(gen, file, g)
	return g
}

// generateFileContent generates the errors definitions, excluding the package statement.
func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}
	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible.")
	g.P("var _ = ", errorsPackage.Ident("New"))
	g.P("var _ = ", contextPackage.Ident("TODO"))
	g.P("var _ = ", asynqPackage.Ident("NewServeMux"))
	g.P("var _ = new(", emptyPackage.Ident("Empty"), ")")
	g.P()

	for _, service := range file.Services {
		genService(gen, file, g, service)
	}
}

func genService(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P("//")
		g.P(deprecationComment)
	}
	// HTTP Server.
	sd := &serviceDesc{
		ServiceType: service.GoName,
		ServiceName: string(service.Desc.FullName()),
		Metadata:    file.Desc.Path(),
	}
	for _, method := range service.Methods {
		if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
			continue
		}
		rule, ok := MatchAsynqRule(method.Comments.Leading)
		if ok {
			sd.Methods = append(sd.Methods, buildAsynqRule(g, method, rule))
		}
	}
	if len(sd.Methods) == 0 {
		return
	}
	err := execute(g, sd)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr,
			"\u001B[31mWARN\u001B[m: generate failed. %v\n", err)
	}
}

func hasHTTPRule(services []*protogen.Service) bool {
	for _, service := range services {
		for _, method := range service.Methods {
			if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
				continue
			}
			if _, ok := MatchAsynqRule(method.Comments.Leading); ok {
				return true
			}
		}
	}
	return false
}

func buildAsynqRule(g *protogen.GeneratedFile, m *protogen.Method, rule *Task) *methodDesc {
	return buildMethodDesc(g, m, rule)
}

func buildMethodDesc(g *protogen.GeneratedFile, m *protogen.Method, rule *Task) *methodDesc {
	defer func() { methodSets[m.GoName]++ }()
	comment := m.Comments.Leading.String() + m.Comments.Trailing.String()
	if comment != "" {
		comment = "// " + m.GoName + strings.TrimPrefix(strings.TrimSuffix(comment, "\n"), "//")
	} else {
		comment = "// " + m.GoName
	}
	return &methodDesc{
		Name:     m.GoName,
		Num:      methodSets[m.GoName],
		Request:  g.QualifiedGoIdent(m.Input.GoIdent),
		Reply:    g.QualifiedGoIdent(m.Output.GoIdent),
		Comment:  comment,
		Pattern:  rule.Pattern,
		CronSpec: rule.CronSpec,
	}
}
