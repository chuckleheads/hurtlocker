#!/bin/bash
# set -x
GOPATH=$(go env GOPATH)

RED='\033[0;31m'
NC='\033[0m' # No Color
printf "\n${RED}Generate Go, Swagger and GRPC Gateway${NC}\n"

shopt -s globstar
for i in components/builder-gateway/api/**/**/; do
  # check if there are proto file is in that directory
  protos=(`find $i -maxdepth 1 -name "*.proto"`)
  if [ ${#protos[@]} -gt 0 ]; then
    # shopt -s nullglob
    list=($i*.proto)
    echo " * $i -> ${list[@]}"

    # service
    protoc -I. \
      -I$GOPATH/src \
      -I$PWD/vendor \
      -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
      --go_out=plugins=grpc:$GOPATH/src \
      ${list[@]}

    # generates grpc gateway
    protoc -I. \
      -I$GOPATH/src \
      -I$PWD/vendor \
      -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
      --grpc-gateway_out=request_context=true,logtostderr=true:$GOPATH/src  \
      ${list[@]}

    # generates swagger output, only generate if a gatway file was generated
    gogw=(`find $i -maxdepth 1 -name "*.gw.go"`)
    if [ ${#gogw[@]} -gt 0 ]; then
      protoc -I. \
        -I$GOPATH/src \
        -I$PWD/vendor \
				-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        --swagger_out=logtostderr=true:$PWD  \
        ${list[@]}
    fi

    # # lint
    # protoc -I. \
    #   -I$GOPATH/src \
    #   -I$PWD/vendor \
    #   -I$PWD/vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    #   --lint_out=$PWD \
    #   $list
  else
    echo " * $i (ignored)"
  fi
done
