//go:generate protoc -I . -I ../../internal/third_party --go_out=paths=source_relative:. asynq.proto
package asynq
