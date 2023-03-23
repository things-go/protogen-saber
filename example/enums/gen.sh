#!/bin/bash

script_dir=$(
    cd $(dirname $0)
    pwd
)                                             # 脚本路径
project_dir=$(dirname $(dirname $script_dir)) # 项目路径

proto_dir=${project_dir}/example/enums
out_dir=${project_dir}/example/enums # 生成代码路径
third_party_dir=${project_dir}/internal/third_party

protoc \
    -I ${proto_dir} \
    -I ${third_party_dir} \
    -I ${project_dir}/protosaber \
    --go_out=${out_dir} \
    --go_opt paths=source_relative \
    --saber-enum_out ${out_dir} \
    --saber-enum_opt paths=source_relative \
    nested.proto \
    non_nested.proto

protoc \
    -I ${proto_dir} \
    -I ${third_party_dir} \
    -I ${project_dir}/protosaber \
    --saber-enum_out ${out_dir} \
    --saber-enum_opt suffix=".example.pb.go" \
    --saber-enum_opt template=${proto_dir}/mapper_template.tpl \
    --saber-enum_opt paths=source_relative \
    --saber-enum_opt merge=true \
    --saber-enum_opt filename=mapper \
    --saber-enum_opt package=enums \
    --saber-enum_opt go_package="github.com/things-go/examples/enums" \
    nested.proto \
    non_nested.proto
