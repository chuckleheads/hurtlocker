# Hurtlocker

gRPC server with built in swagger support. This is far from being anywhere near functional

## Making changes

#### Generating go code from protobufs
*** you need to do this for each proto file you change.
*** TODO: implement a makefile for this
```
 protoc -I/usr/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --go_out=plugins=grpc:$GOPATH/src  --grpc-gateway_out=request_context=true,logtostderr=true:$GOPATH/src components/gateway/api/origins/origins.proto
 ```

#### Generating swagger defs
```
protoc -I/usr/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --swagger_out=logtostderr=true:. components/gateway/api/origins/origins.proto
```
