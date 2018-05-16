.DEFAULT_GOAL := all

all: protobufs

# Order Matters

protobufs:
	components/originsrv/scripts/grpc.sh
	components/gateway/scripts/grpc.sh
