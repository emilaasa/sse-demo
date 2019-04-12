PROTO_PATH  = proto/simple.proto

.PHONY: proto
proto: 
	protoc --go_out=plugins=grpc:. \
		${PROTO_PATH}

