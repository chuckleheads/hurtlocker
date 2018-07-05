.DEFAULT_GOAL := all

all: worker

# Order Matters

originsrv:
	components/originsrv/scripts/grpc.sh

agent: originsrv
	components/agent/scripts/grpc.sh

sessionsrv: agent
	components/sessionsrv/scripts/grpc.sh

logsrv: sessionsrv
	components/logsrv/scripts/grpc.sh

worker: logsrv
	components/worker/scripts/grpc.sh
