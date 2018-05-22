package server

import (
	"context"
	"database/sql"

	"github.com/chuckleheads/hurtlocker/components/sessionsrv/accounts/request"
	"github.com/chuckleheads/hurtlocker/components/sessionsrv/accounts/response"
	"github.com/chuckleheads/hurtlocker/components/sessionsrv/data_store/functions"
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

func (srv *Server) CreateAccount(ctx context.Context, req *request.CreateAccountReq) (*response.AccountResp, error) {
	var id int64
	// JB: This should be a postgres function, not inline code
	err := srv.db.
		QueryRow(functions.InsertAccountV1, req.Username, req.Email).
		Scan(&id)

	// JB: Handle other errors here like row conflicts
	if err != nil {
		panic(err.Error())
	}

	account := response.Account{
		Id:       id,
		Username: req.Username,
		Email:    req.Email,
	}

	accountResp := response.AccountResp{
		Account: &account,
	}

	return &accountResp, nil
}

func (srv *Server) GetAccount(ctx context.Context, req *request.GetAccountReq) (*response.AccountResp, error) {
	var account response.Account

	err := srv.db.
		QueryRow(functions.GetAccountByUsernameV1, req.Username).
		Scan(&account.Id, &account.Username, &account.Email)

	if err != nil && err != sql.ErrNoRows {
		panic(err.Error())
	}

	if err == nil {
		resp := response.AccountResp{
			Account: &account,
		}

		return &resp, nil
	} else {
		return nil, grpc.Errorf(codes.NotFound, "%s", req.Username)
	}
}

func (srv *Server) UpdateAccount(ctx context.Context, req *request.UpdateAccountReq) (*response.AccountResp, error) {
	var account response.Account
	err := srv.db.
		QueryRow(functions.UpdateAccountV1, req.Email, req.Username).
		Scan(&account.Id, &account.Username, &account.Email)

	if err != nil {
		panic(err.Error())
	}

	actresp := response.AccountResp{
		Account: &account,
	}

	return &actresp, nil
}
