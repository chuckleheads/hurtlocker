package server

import (
	"context"
	"database/sql"

	"github.com/chuckleheads/hurtlocker/components/originsrv/data_store/functions"
	"github.com/chuckleheads/hurtlocker/components/originsrv/projects"
	"github.com/chuckleheads/hurtlocker/components/originsrv/projects/request"
	"github.com/chuckleheads/hurtlocker/components/originsrv/projects/response"
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

func (srv *Server) CreateProject(ctx context.Context, req *request.CreateProjectReq) (*response.ProjectResp, error) {
	var project projects.Project
	p := req.Project

	err := srv.db.
		QueryRow(functions.InsertProjectV1, p.Origin, p.PackageName, p.PlanPath, p.VcsType, p.VcsData).
		Scan(&project.Id, &project.Origin, &project.PackageName, &project.PlanPath, &project.VcsType, &project.VcsData)

	// JB TODO: Handle other errors here like row conflicts
	if err != nil {
		panic(err.Error())
	}

	resp := response.ProjectResp{
		Project: &project,
	}

	return &resp, nil
}

func (srv *Server) GetProject(ctx context.Context, req *request.GetProjectReq) (*response.ProjectResp, error) {
	var project projects.Project

	err := srv.db.
		QueryRow(functions.GetProjectV1, req.Origin, req.PackageName).
		Scan(&project.Id, &project.Origin, &project.PackageName, &project.PlanPath, &project.VcsType, &project.VcsData)

	if err != nil && err != sql.ErrNoRows {
		panic(err.Error())
	}

	if err == nil {
		resp := response.ProjectResp{
			Project: &project,
		}

		return &resp, nil
	} else {
		return nil, grpc.Errorf(codes.NotFound, "%s/%s", req.Origin, req.PackageName)
	}
}

func (srv *Server) UpdateProject(ctx context.Context, req *request.UpdateProjectReq) (*response.ProjectResp, error) {
	var project projects.Project
	p := req.Project

	err := srv.db.
		QueryRow(functions.UpdateProjectV1, p.Origin, p.PackageName, p.PlanPath, p.VcsType, p.VcsData).
		Scan(&project.Id, &project.Origin, &project.PackageName, &project.PlanPath, &project.VcsType, &project.VcsData)

	if err != nil {
		panic(err.Error())
	}

	resp := response.ProjectResp{
		Project: &project,
	}

	return &resp, nil
}
