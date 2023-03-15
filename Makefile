generate: 
	@protoc -I . -I ./internal/third_party --go_out=paths=source_relative:. protosaber/asynq/asynq.proto  protosaber/enumerate/enumerate.proto

.PHONY: generate