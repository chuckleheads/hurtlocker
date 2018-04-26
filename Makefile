.DEFAULT_GOAL := all

all: protobufs swagger

protobufs:
	protoc -I/usr/include -I. -I$(GOPATH)/src -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:$(GOPATH)/src --grpc-gateway_out=request_context=true,logtostderr=true:$(GOPATH)/src components/gateway/api/origins/origins.proto

swagger:
	protoc -I/usr/include -I. -I$(GOPATH)/src -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. components/gateway/api/origins/origins.proto
