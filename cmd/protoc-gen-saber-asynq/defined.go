package main

import (
	"strings"

	"github.com/things-go/protogen-saber/internal/protoutil"
	"google.golang.org/protobuf/compiler/protogen"
)

const version = "v0.3.0"

// annotation const value
const (
	annotation_Path         = "asynq"
	annotation_Key_Pattern  = "pattern"
	annotation_Key_CronSpec = "cron_spec"
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
	Pattern  string // 匹配器
	CronSpec string // cron specification
}

type Task struct {
	Pattern  string
	CronSpec string
}

func MatchAsynqRule(c protogen.Comments) (*Task, bool) {
	annotes := protoutil.NewComments(c).FindAnnotation(annotation_Path)
	if len(annotes) > 0 {
		t := &Task{}
		for _, v := range annotes {
			if strings.EqualFold(v.Key, annotation_Key_Pattern) {
				t.Pattern = v.Value
			} else if strings.EqualFold(v.Key, annotation_Key_CronSpec) {
				t.CronSpec = v.Value
			}
		}
		if t.Pattern != "" && t.CronSpec != "" {
			return t, true
		}
	}
	return nil, false
}
