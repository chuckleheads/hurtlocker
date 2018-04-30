package server

import (
	"context"
	"fmt"

	"github.com/chuckleheads/hurtlocker/components/originsrv/origins/request"
	"github.com/chuckleheads/hurtlocker/components/originsrv/origins/response"
)

// TODO should be a server handle
type Server struct {
	db string
}

func NewServer(db string) *Server {
	return &Server{db}
}

func (srv *Server) CreateOrigin(ctx context.Context, req *request.CreateOriginReq) (*response.OriginResp, error) {
	fmt.Printf("HERE")
	origin := response.Origin{
		Id:   "a thing?",
		Name: "foo",
		DefaultPackageVisibility: "shrug",
	}
	orgresp := response.OriginResp{
		Origin: &origin,
	}
	return &orgresp, nil
}

func (srv *Server) GetOrigin(ctx context.Context, req *request.GetOriginReq) (*response.OriginResp, error) {
	origin := response.Origin{
		Id:   "a thing?",
		Name: "foo",
		DefaultPackageVisibility: "shrug",
	}
	orgresp := response.OriginResp{
		Origin: &origin,
	}
	return &orgresp, nil
}

func (srv *Server) UpdateOrigin(ctx context.Context, req *request.UpdateOriginReq) (*response.OriginResp, error) {
	origin := response.Origin{
		Id:   "a thing?",
		Name: "foo",
		DefaultPackageVisibility: "shrug",
	}
	orgresp := response.OriginResp{
		Origin: &origin,
	}
	return &orgresp, nil
}
