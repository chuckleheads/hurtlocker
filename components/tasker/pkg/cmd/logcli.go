package cmd

import (
	pb "github.com/chuckleheads/hurtlocker/components/logsrv/logrecv"

	"google.golang.org/grpc"
)

func LogSrvClient() pb.LogRecvClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure(), grpc.WithAuthority("logsrv"))
	conn, err := grpc.Dial(":9211", opts...)
	if err != nil {
		panic(err.Error())
	}
	return pb.NewLogRecvClient(conn)
}
