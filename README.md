# go-gRPC
書籍スターティングgRPCの写経を行うリポジトリ

mkdir -p app/gen/api/pancake/maker

bundle exec grpc_tools_ruby_protoc \
    -I ../proto \
    --ruby_out=app/gen/api/pancake/maker \
    --grpc_out=app/gen/api/pancake/maker \
    ../proto/pancake.proto
