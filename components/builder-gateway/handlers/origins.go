package handlers

import (
	"golang.org/x/net/context"
	"log"

	pbreq "github.com/chuckleheads/hurtlocker/components/builder-gateway/api/origins/request"
	pbres "github.com/chuckleheads/hurtlocker/components/builder-gateway/api/origins/response"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type server struct {
	origins map[string]pbres.Origin
}

func NewOriginsServer() server {
	return server{
		origins: make(map[string]pbres.Origin),
	}
}

func (s server) GetOrigin(ctx context.Context, req *pbreq.GetOriginReq) (*pbres.OriginResp, error) {
	log.Println("Getting origin!")

	// What should happen here is use the client that is defined in the "OriginServer" struct above
	// which will allow you to talk gRPC to the "origin service" that will actually do database things

	if req.Name == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "Origin Name cannot be empty")
	}

	origin, exists := s.origins[req.Name]

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

	s.origins[req.Name] = pbres.Origin{
		Id:   id,
		Name: req.Name,
		DefaultPackageVisibility: req.DefaultPackageVisibility,
	}

	origin := s.origins[req.Name]

	ores := pbres.OriginResp{
		Origin: &origin,
	}

	log.Println("Origin created!")
	return &ores, nil
}

func (s server) UpdateOrigin(ctx context.Context, req *pbreq.UpdateOriginReq) (*pbres.OriginResp, error) {
	log.Println("Updating origin!")

	if req.Name == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "Origin Name cannot be empty")
	}

	origin := s.origins[req.Name]
	origin.DefaultPackageVisibility = req.DefaultPackageVisibility
	s.origins[req.Name] = origin

	ores := pbres.OriginResp{
		Origin: &origin,
	}

	log.Println("Origin updated!")
	return &ores, nil
}
