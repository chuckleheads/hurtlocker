.DEFAULT_GOAL := all

all: protobufs

# Order Matters

protobufs:
	components/originsrv/scripts/grpc.sh
	components/builder-gateway/scripts/grpc.sh
