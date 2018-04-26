package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/google/uuid"

	pb "github.com/chuckleheads/hurtlocker/components/builder-gateway/api/origins"
	pbreq "github.com/chuckleheads/hurtlocker/components/builder-gateway/api/origins/request"
	pbres "github.com/chuckleheads/hurtlocker/components/builder-gateway/api/origins/response"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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
	pb.RegisterOriginsServer(server, OriginsServer())

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

type server struct {
	origins map[string]pbres.Origin
}

func OriginsServer() server {
	return server{
		origins: make(map[string]pbres.Origin),
	}
}

func (s server) GetOrigin(ctx context.Context, req *pbreq.GetOriginReq) (*pbres.OriginResp, error) {
	log.Println("Getting origin!")

	if req.Id == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "Origin ID cannot be empty")
	}

	origin, exists := s.origins[req.Id]
	if !exists {
		return nil, grpc.Errorf(codes.NotFound, "origin not found")
	}

	ores := pbres.OriginResp{
		Origin: &origin,
	}

	log.Println("Origin found!")
	return &ores, nil
}

func (s server) CreateOrigin(ctx context.Context, req *pbreq.CreateOriginReq) (*pbres.OriginResp, error) {
	log.Println("Creating origin!")

	if req.Name == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "Origin Name cannot be empty")
	}

	id := uuid.New().String()

	s.origins[id] = pbres.Origin{
		Id:   id,
		Name: req.Name,
		DefaultPackageVisibility: req.DefaultPackageVisibility,
	}

	origin := s.origins[id]

	ores := pbres.OriginResp{
		Origin: &origin,
	}

	log.Println("Origin created!")
	return &ores, nil
}

func (s server) UpdateOrigin(ctx context.Context, req *pbreq.UpdateOriginReq) (*pbres.OriginResp, error) {
	log.Println("Updating origin!")

	if req.Id == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "Origin Id cannot be empty")
	}

	origin := s.origins[req.Id]
	origin.DefaultPackageVisibility = req.DefaultPackageVisibility
	s.origins[req.Id] = origin

	ores := pbres.OriginResp{
		Origin: &origin,
	}

	log.Println("Origin updated!")
	return &ores, nil
}
