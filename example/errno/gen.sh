#!/bin/bash

script_dir=$(
    cd $(dirname $0)
    pwd
)                                             # 脚本路径
project_dir=$(dirname $(dirname $script_dir)) # 项目路径

proto_dir=${project_dir}/example/errno
out_dir=${project_dir}/example/errno # 生成代码路径
third_party_dir=${project_dir}/internal/third_party

protoc \
    -I ${proto_dir} \
    -I ${third_party_dir} \
    --go_out=${out_dir} \
    --go_opt paths=source_relative \
    --saber-errno_out ${out_dir} \
    --saber-errno_opt paths=source_relative \
    --saber-errno_opt template=builtin-est \
    errno.proto
