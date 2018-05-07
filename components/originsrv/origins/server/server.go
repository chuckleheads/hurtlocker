package server

import (
	"context"
	"database/sql"

	"github.com/chuckleheads/hurtlocker/components/originsrv/data_store"
	"github.com/chuckleheads/hurtlocker/components/originsrv/origins/request"
	"github.com/chuckleheads/hurtlocker/components/originsrv/origins/response"
	_ "github.com/lib/pq"
)

// TODO should be a server handle
type Server struct {
	db sql.DB
}

func NewServer(dbConfig data_store.DBConfig) *Server {
	return &Server{db: *data_store.New(&dbConfig)}
}

func (srv *Server) CreateOrigin(ctx context.Context, req *request.CreateOriginReq) (*response.OriginResp, error) {
	res, err := srv.db.Exec("INSERT INTO origins(name, default_package_visibility) VALUES(?,?)", req.Name, req.DefaultPackageVisibility)
	if err == nil {
		panic(err.Error())
	}
	id, err := res.LastInsertId()
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
	origin := response.Origin{
		Id:   12345,
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
		Id:   12345,
		Name: "foo",
		DefaultPackageVisibility: "shrug",
	}
	orgresp := response.OriginResp{
		Origin: &origin,
	}
	return &orgresp, nil
}
