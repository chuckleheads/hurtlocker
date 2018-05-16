package handlers

import (
	"log"

	"golang.org/x/net/context"

	pbreq "github.com/chuckleheads/hurtlocker/components/originsrv/origins/request"
	pbres "github.com/chuckleheads/hurtlocker/components/originsrv/origins/response"

	"github.com/chuckleheads/hurtlocker/components/originsrv/origins"

	"google.golang.org/grpc"
)

type OriginSrv struct {
	client origins.OriginsClient
}

func NewOriginsServer(client origins.OriginsClient) *OriginSrv {
	return &OriginSrv{
		client: client,
	}
}

func Dialer() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:7001", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	// defer conn.Close()
	return conn
}

func (s OriginSrv) GetOrigin(ctx context.Context, req *pbreq.GetOriginReq) (*pbres.OriginResp, error) {
	log.Println("Getting origin!")

	// These need to change to only get passed what they need - too tired to think about it
	return s.client.GetOrigin(ctx, req)
}

func (s OriginSrv) CreateOrigin(ctx context.Context, req *pbreq.CreateOriginReq) (*pbres.OriginResp, error) {
	log.Println("Creating origin!")

	return s.client.CreateOrigin(ctx, req)
}

func (s OriginSrv) UpdateOrigin(ctx context.Context, req *pbreq.UpdateOriginReq) (*pbres.OriginResp, error) {
	log.Println("Updating origin!")

	return s.client.UpdateOrigin(ctx, req)
}
