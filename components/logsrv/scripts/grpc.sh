#!/bin/bash
# set -x
GOPATH=$(go env GOPATH)

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color
printf "\n${GREEN}Generate Go, Swagger and GRPC Gateway${NC}\n"

shopt -s globstar
for i in components/logsrv/**/**/; do
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
