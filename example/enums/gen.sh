scriptDir=$(
    cd $(dirname $0)
    pwd
)                                        # 脚本路径
projDir=$(dirname $(dirname $scriptDir)) # 项目路径

protoDir=${projDir}/example/enums
outDir=${projDir}/example/enums # 生成代码路径
thirdPartyDir=${projDir}/internal/third_party

protoc \
    -I ${protoDir} \
    -I ${thirdPartyDir} \
    -I ${projDir}/protosaber \
    --go_out=${outDir} \
    --go_opt paths=source_relative \
    --saber-enum_out ${outDir} \
    --saber-enum_opt paths=source_relative \
    nested.proto \
    non_nested.proto

protoc \
    -I ${protoDir} \
    -I ${thirdPartyDir} \
    -I ${projDir}/protosaber \
    --saber-enum_out ${outDir} \
    --saber-enum_opt suffix=".example.pb.go" \
    --saber-enum_opt template=${protoDir}/mapper_template.tpl \
    --saber-enum_opt paths=source_relative \
    --saber-enum_opt merge=true \
    --saber-enum_opt filename=mapper \
    --saber-enum_opt package=enums \
    --saber-enum_opt go_package="github.com/things-go/examples/enums" \
    nested.proto \
    non_nested.proto
