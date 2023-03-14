scriptDir=$(
    cd $(dirname $0)
    pwd
)                                        # 脚本路径
projDir=$(dirname $(dirname $scriptDir)) # 项目路径

protoDir=${projDir}/example/asynq
outDir=${projDir}/example/asynq # 生成代码路径
thirdPartyDir=${projDir}/internal/third_party

protoc \
    -I ${protoDir} \
    -I ${thirdPartyDir} \
    -I ${projDir}/protosaber \
    --go_out=${outDir} \
    --go_opt paths=source_relative \
    --saber-asynq_out ${outDir} \
    --saber-asynq_opt paths=source_relative \
    asynq.proto
