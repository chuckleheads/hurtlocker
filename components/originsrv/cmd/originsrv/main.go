package main

import (
	"log"
	"net"

	pb "github.com/chuckleheads/hurtlocker/components/originsrv/origins"
	srv "github.com/chuckleheads/hurtlocker/components/originsrv/origins/server"
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
	db := server.EstablishConnection()
	server.RunMigrations(db)

	pb.RegisterOriginsServer(server, srv.NewServer("some_pg_server"))

	log.Printf("gRPC Listening on %s\n", lis.Addr().String())
	server.Serve(lis)
}
