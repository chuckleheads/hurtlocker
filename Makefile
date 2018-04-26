.DEFAULT_GOAL := all

all: protobufs

protobufs:
	components/builder-gateway/scripts/grpc.sh
