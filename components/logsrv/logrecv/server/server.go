package server

import (
	"fmt"
	"io"
	"log"

	pb "github.com/chuckleheads/hurtlocker/components/logsrv/logrecv"
	pbr "github.com/chuckleheads/hurtlocker/components/logsrv/logrecv/response"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (srv *Server) ReceiveLogs(stream pb.LogRecv_ReceiveLogsServer) error {
	fmt.Println("Logging...")
	for {
		line, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pbr.LogSummary{
				Status: true,
			})
		}
		if err != nil {
			return err
		}
		log.Println(line.GetStderrLine())
		log.Println(line.GetStdoutLine())
	}
}
