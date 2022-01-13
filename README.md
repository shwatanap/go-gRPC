# go-gRPC
書籍スターティングgRPCの写経を行うリポジトリ

.protoファイルから自動生成させるコマンド
bundle exec grpc_tools_ruby_protoc \
    -I ../proto \
    --ruby_out=app/gen/api/pancake/baker \
    --grpc_out=app/gen/api/pancake/baker \
    ../proto/pancake.proto
