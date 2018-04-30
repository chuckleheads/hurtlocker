package server

import "context"

// TODO should be a server handle
type Server struct {
	db string
}

func New(db string) *Server {
	return &Server{db}
}

func (srv *Server) Create(ctx context.Context) {

}

func (srv *Server) Fetch(ctx context.Context) {

}

func (srv *Server) Update(ctx context.Context) {

}
