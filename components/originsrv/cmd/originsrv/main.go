package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	addr := ":7000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to initializa TCP listen: %v", err)
	}
	defer lis.Close()

	runGRPC(lis)
}

func runGRPC(lis net.Listener) {
	server := grpc.NewServer()
	// pb.RegisterOriginsServer(server, handlers.NewOriginsServer())

	log.Printf("gRPC Listening on %s\n", lis.Addr().String())
	server.Serve(lis)
}
