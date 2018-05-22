package server

import (
	"context"
	"database/sql"

	"github.com/chuckleheads/hurtlocker/components/originsrv/data_store/functions"
	"github.com/chuckleheads/hurtlocker/components/originsrv/origins/request"
	"github.com/chuckleheads/hurtlocker/components/originsrv/origins/response"
	_ "github.com/lib/pq"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Server struct {
	db *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{db}
}

func (srv *Server) CreateOrigin(ctx context.Context, req *request.CreateOriginReq) (*response.OriginResp, error) {
	var id int64
	err := srv.db.
		QueryRow(functions.InsertOriginV1, req.Name, req.DefaultPackageVisibility).
		Scan(&id)

	// TED: Handle other errors here like row conflicts
	if err != nil {
		panic(err.Error())
	}

	origin := response.Origin{
		Id:   id,
		Name: req.Name,
		DefaultPackageVisibility: req.DefaultPackageVisibility,
	}

	originResp := response.OriginResp{
		Origin: &origin,
	}

	return &originResp, nil
}

func (srv *Server) GetOrigin(ctx context.Context, req *request.GetOriginReq) (*response.OriginResp, error) {
	var origin response.Origin

	err := srv.db.
		QueryRow(functions.GetOriginByNameV1, req.Name).
		Scan(&origin.Id, &origin.Name, &origin.DefaultPackageVisibility)

	if err != nil && err != sql.ErrNoRows {
		panic(err.Error())
	}

	if err == nil {
		resp := response.OriginResp{
			Origin: &origin,
		}

		return &resp, nil
	} else {
		return nil, grpc.Errorf(codes.NotFound, "%s", req.Name)
	}
}

func (srv *Server) UpdateOrigin(ctx context.Context, req *request.UpdateOriginReq) (*response.OriginResp, error) {
	var origin response.Origin
	err := srv.db.
		QueryRow(functions.UpdateOriginV1, req.DefaultPackageVisibility, req.Name).
		Scan(&origin.Id, &origin.Name, &origin.DefaultPackageVisibility)

	if err != nil {
		panic(err.Error())
	}

	orgresp := response.OriginResp{
		Origin: &origin,
	}

	return &orgresp, nil
}
