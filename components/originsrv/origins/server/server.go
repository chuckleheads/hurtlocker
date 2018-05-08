package server

import (
	"context"
	"database/sql"

	"github.com/chuckleheads/hurtlocker/components/originsrv/data_store"
	"github.com/chuckleheads/hurtlocker/components/originsrv/origins/request"
	"github.com/chuckleheads/hurtlocker/components/originsrv/origins/response"
	_ "github.com/lib/pq"
)

type Server struct {
	db sql.DB
}

func NewServer(dbConfig data_store.DBConfig) *Server {
	return &Server{db: *data_store.New(&dbConfig)}
}

func (srv *Server) CreateOrigin(ctx context.Context, req *request.CreateOriginReq) (*response.OriginResp, error) {
	var id int64
	// TED: This should be a postgres function, not inline code
	err := srv.db.
		QueryRow("INSERT INTO origins(name, default_package_visibility) VALUES($1,$2) RETURNING id", req.Name, req.DefaultPackageVisibility).
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
		QueryRow("Select * from origins where name = $1", req.Name).
		Scan(&origin.Id, &origin.Name, &origin.DefaultPackageVisibility)
	if err != nil {
		panic(err.Error())
	}
	orgresp := response.OriginResp{
		Origin: &origin,
	}
	return &orgresp, nil
}

func (srv *Server) UpdateOrigin(ctx context.Context, req *request.UpdateOriginReq) (*response.OriginResp, error) {
	var origin response.Origin
	err := srv.db.
		QueryRow("UPDATE origins set default_package_visibility = $1 where name = $2 RETURNING *", req.DefaultPackageVisibility, req.Name).
		Scan(&origin.Id, &origin.Name, &origin.DefaultPackageVisibility)
	if err != nil {
		panic(err.Error())
	}
	orgresp := response.OriginResp{
		Origin: &origin,
	}
	return &orgresp, nil
}
