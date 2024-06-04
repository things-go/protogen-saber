package main

import (
	"github.com/things-go/proc/proc"
	"google.golang.org/protobuf/compiler/protogen"
)

// annotation const value
const (
	Identity                = "asynq"
	Attribute_Name_Pattern  = "pattern"
	Attribute_Name_CronSpec = "cron_spec"
)

type Task struct {
	Enabled  bool
	Pattern  string
	CronSpec string
}

func IsDeriveTaskEnabled(s protogen.Comments) bool {
	derives, _ := proc.NewCommentLines(string(s)).FindDerives(Identity)
	return proc.Derives(derives).ContainHeadless(Identity)
}

func ParserDeriveTask(s protogen.Comments) *Task {
	ret := &Task{}
	derives, _ := proc.NewCommentLines(string(s)).FindDerives(Identity)
	for _, annotate := range derives {
		if annotate.Headless() {
			ret.Enabled = true
			continue
		}
		for _, attr := range annotate.Attrs {
			switch attr.Name {
			case Attribute_Name_Pattern:
				if vv, ok := attr.Value.(proc.String); ok {
					ret.Pattern = vv.Value
				}
			case Attribute_Name_CronSpec:
				if vv, ok := attr.Value.(proc.String); ok {
					ret.CronSpec = vv.Value
				}
			}
		}
	}
	return ret
}
