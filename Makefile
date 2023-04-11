generate: 
	@protoc -I . \
	-I ./internal/third_party \
	--go_out=paths=source_relative:. \
	protosaber/asynq/asynq.proto  \
	protosaber/enumerate/enumerate.proto \
	protosaber/seaql/seaql.proto

tools:
	go install github.com/things-go/protogen-saber/cmd/protoc-gen-saber-asynq
	go install github.com/things-go/protogen-saber/cmd/protoc-gen-saber-enum
	go install github.com/things-go/protogen-saber/cmd/protoc-gen-saber-seaql

.PHONY: generate