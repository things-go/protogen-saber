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
	go install github.com/things-go/protogen-saber/cmd/protoc-gen-saber-model
	go install github.com/things-go/protogen-saber/cmd/protoc-gen-saber-rapier

gen.asynq:
	@./example/asynq/gen.sh
gen.enum:
	@./example/enums/gen.sh
gen.seaql:
	@./example/seaql/gen.sh

.PHONY: generate tools