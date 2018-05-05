package server

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/net/context"

	"github.com/chuckleheads/hurtlocker/components/originsrv/origins/request"
	"github.com/chuckleheads/hurtlocker/components/originsrv/origins/response"
	_ "github.com/lib/pq"
)

// TODO should be a server handle
type Server struct {
	db string
}

func NewServer(db string) *Server {
	return &Server{db}
}

// This might not be the right place for these two methods, but somehow we need to share the db connection between all these methods. It'd be rad if there was a way to shove it into the Context that all the handlers get. Or maybe there's a more idiomatic way?
func (srv *Server) EstablishConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://hab@localhost:26257/hurtlocker?sslmode=disable")

	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	return db, err
}

func (srv *Server) RunMigrations(db *sql.DB) {
	// This format is not ideal =/
	log.Println("Running db migrations")
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS origins (id bigserial PRIMARY KEY, name TEXT NOT NULL UNIQUE, default_package_visibility text NOT NULL DEFAULT 'public')"); err != nil {
		log.Fatal(err)
	}
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
