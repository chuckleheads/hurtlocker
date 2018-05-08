package main

import (
	"log"
	"net"

	"github.com/chuckleheads/hurtlocker/components/originsrv/data_store"
	pb "github.com/chuckleheads/hurtlocker/components/originsrv/origins"
	srv "github.com/chuckleheads/hurtlocker/components/originsrv/origins/server"
	"google.golang.org/grpc"
)

func main() {
	addr := ":7001"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to initializa TCP listen: %v", err)
	}
	defer lis.Close()

	runGRPC(lis)
}

func runGRPC(lis net.Listener) {
	server := grpc.NewServer()

	pgCreds := data_store.DBConfig{
		Port:     26257,
		Host:     "localhost",
		SSLMode:  "disable",
		Database: "originsrv",
		Username: "root",
		Password: "",
	}
	db := data_store.New(&pgCreds)
	pb.RegisterOriginsServer(server, srv.NewServer(db))

	log.Printf("gRPC Listening on %s\n", lis.Addr().String())
	server.Serve(lis)
}
