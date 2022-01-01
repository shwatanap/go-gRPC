.PHONY: setup
setup:
	go get -u github.com/golang/protobuf/protoc-gen-go

.PHONY: protoc
protoc:
	protoc \
		-Iproto \
		--go_out=plugins=grpc:api \
		proto/*.proto

.PHONY: run
run:
	cd grpc && fresh

ls:
	grpc_cli ls localhost:50051 pancake.baker.PancakeBakerService -l

call:
	grpc_cli call localhost:50051 pancake.baker.PancakeBakerService.Bake 'menu:1'
