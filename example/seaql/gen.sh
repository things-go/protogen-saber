#!/bin/bash

script_dir=$(
    cd $(dirname $0)
    pwd
)                                             # 脚本路径
project_dir=$(dirname $(dirname $script_dir)) # 项目路径

proto_dir=${project_dir}/example/seaql
out_dir=${project_dir}/example/seaql # 生成代码路径
third_party_dir=${project_dir}/internal/third_party

protoc \
    -I ${proto_dir} \
    -I ${third_party_dir} \
    --saber-seaql_out ${out_dir} \
    --saber-seaql_opt paths=source_relative \
    --saber-seaql_opt trim_prefix=false \
    seaql.proto

protoc \
    -I ${proto_dir} \
    -I ${third_party_dir} \
    --saber-seaql_out ${out_dir} \
    --saber-seaql_opt paths=source_relative \
    --saber-seaql_opt merge=true \
    seaql.proto

# protoc \
#     -I ${proto_dir} \
#     -I ${third_party_dir} \
#     --saber-model_out ${out_dir}/model \
#     --saber-model_opt paths=source_relative \
#     --saber-model_opt package=model \
#     seaql.proto

# protoc \
#     -I ${proto_dir} \
#     -I ${third_party_dir} \
#     --saber-assist_out ${out_dir}/repo \
#     --saber-assist_opt paths=source_relative \
#     --saber-assist_opt package=model \
#     --saber-assist_opt model_source=github.com/proa/model \
#     seaql.proto
