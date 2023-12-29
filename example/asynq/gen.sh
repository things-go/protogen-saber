#!/bin/bash

script_dir=$(
    cd $(dirname $0)
    pwd
)                                             # 脚本路径
project_dir=$(dirname $(dirname $script_dir)) # 项目路径

proto_dir=${project_dir}/example/asynq
out_dir=${project_dir}/example/asynq # 生成代码路径
third_party_dir=${project_dir}/internal/third_party

protoc \
    -I ${proto_dir} \
    -I ${third_party_dir} \
    --go_out=${out_dir} \
    --go_opt paths=source_relative \
    --saber-asynq_out ${out_dir} \
    --saber-asynq_opt paths=source_relative \
    --saber-asynq_opt disable_saber=true \
    --saber-asynq_opt codec=json \
    asynq.proto
