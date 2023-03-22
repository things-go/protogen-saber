scriptDir=$(
    cd $(dirname $0)
    pwd
)                                        # 脚本路径
projDir=$(dirname $(dirname $scriptDir)) # 项目路径

protoDir=${projDir}/example/seaql
outDir=${projDir}/example/seaql # 生成代码路径
thirdPartyDir=${projDir}/internal/third_party

protoc \
    -I ${protoDir} \
    -I ${thirdPartyDir} \
    -I ${projDir}/protosaber \
    --saber-seaql_out ${outDir} \
    --saber-seaql_opt paths=source_relative \
    seaql.proto
