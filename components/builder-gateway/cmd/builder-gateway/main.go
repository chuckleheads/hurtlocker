package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"golang.org/x/net/context"

	pb "github.com/chuckleheads/hurtlocker/components/builder-gateway/api/origins"
	"github.com/chuckleheads/hurtlocker/components/builder-gateway/handlers"
	originsrv "github.com/chuckleheads/hurtlocker/components/originsrv/origins"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {
	addr := ":6000"
	clientAddr := fmt.Sprintf("localhost%s", addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to initializa TCP listen: %v", err)
	}
	defer lis.Close()

	go runGRPC(lis)
	runHTTP(clientAddr)
}

func runGRPC(lis net.Listener) {
	server := grpc.NewServer()
	pb.RegisterOriginsServer(server, handlers.NewOriginsServer(originsrv.NewOriginsClient(handlers.Dialer())))

	log.Printf("gRPC Listening on %s\n", lis.Addr().String())
	server.Serve(lis)
}

func runHTTP(clientAddr string) {
	addr := ":6001"
	opts := []grpc.DialOption{grpc.WithInsecure()}
	mux := runtime.NewServeMux()

	if err := pb.RegisterOriginsHandlerFromEndpoint(context.Background(), mux, clientAddr, opts); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}

	log.Printf("HTTP Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}