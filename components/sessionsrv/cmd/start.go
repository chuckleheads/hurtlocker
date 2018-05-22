package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/chuckleheads/hurtlocker/components/sessionsrv/accounts"
	srv "github.com/chuckleheads/hurtlocker/components/sessionsrv/accounts/server"
	pba "github.com/chuckleheads/hurtlocker/components/sessionsrv/api/accounts"
	"github.com/chuckleheads/hurtlocker/components/sessionsrv/config"
	"github.com/chuckleheads/hurtlocker/components/sessionsrv/data_store"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Session Server is rad",
	Long:  `Wouldn't it be nice to have a long description here.`,
	Run: func(cmd *cobra.Command, args []string) {
		setup()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func setup() {
	config, err := ConfigFromViper()
	if err != nil {
		panic(err.Error())
	}
	addr := fmt.Sprintf(":%d", config.Port)
	clientAddr := fmt.Sprintf("localhost%s", addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to initializa TCP listen: %v", err)
	}
	defer lis.Close()

	go runGRPC(config, lis)
	runHTTP(config, clientAddr)
}

func runGRPC(config *config.Config, lis net.Listener) {
	server := grpc.NewServer()
	db := data_store.New(&config.Datastore)
	pb.RegisterAccountsServer(server, srv.NewServer(db))

	log.Printf("gRPC Listening on %s\n", lis.Addr().String())
	server.Serve(lis)
}

func runHTTP(config *config.Config, clientAddr string) {
	addr := fmt.Sprintf(":%d", config.HttpPort)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	mux := runtime.NewServeMux()

	if err := pba.RegisterAccountsHandlerFromEndpoint(context.Background(), mux, clientAddr, opts); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}

	log.Printf("HTTP Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
