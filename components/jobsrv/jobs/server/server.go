package server

import (
	"context"
	"database/sql"

	"github.com/chuckleheads/hurtlocker/components/jobsrv/data_store/functions"
	"github.com/chuckleheads/hurtlocker/components/jobsrv/jobs/request"
	"github.com/chuckleheads/hurtlocker/components/jobsrv/jobs/response"
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

func (srv *Server) CreateJob(ctx context.Context, req *request.CreateJobReq) (*response.JobResp, error) {
	var id int64
	err := srv.db.
		QueryRow(functions.InsertJobV1, req.Name, req.DefaultPackageVisibility).
		Scan(&id)

	// TED: Handle other errors here like row conflicts
	if err != nil {
		panic(err.Error())
	}

	job := response.Job{
		Id:   id,
		Name: req.Name,
		DefaultPackageVisibility: req.DefaultPackageVisibility,
	}

	jobResp := response.JobResp{
		Job: &job,
	}

	return &jobResp, nil
}

func (srv *Server) GetJob(ctx context.Context, req *request.GetJobReq) (*response.JobResp, error) {
	var job response.Job

	err := srv.db.
		QueryRow(functions.GetJobByNameV1, req.Name).
		Scan(&job.Id, &job.Name, &job.DefaultPackageVisibility)

	if err != nil && err != sql.ErrNoRows {
		panic(err.Error())
	}

	if err == nil {
		resp := response.JobResp{
			Job: &job,
		}

		return &resp, nil
	} else {
		return nil, grpc.Errorf(codes.NotFound, "%s", req.Name)
	}
}
