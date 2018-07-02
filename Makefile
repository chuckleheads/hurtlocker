.DEFAULT_GOAL := all

all: protobufs

# Order Matters

protobufs:
	components/originsrv/scripts/grpc.sh
	components/agent/scripts/grpc.sh
	components/sessionsrv/scripts/grpc.sh
	components/logsrv/scripts/grpc.sh
	components/worker/scripts/grpc.sh
