package cmd

import (
	"log"
	"net"

	pb "github.com/chuckleheads/hurtlocker/components/sessionsrv/accounts"
	srv "github.com/chuckleheads/hurtlocker/components/sessionsrv/accounts/server"
	"github.com/chuckleheads/hurtlocker/components/sessionsrv/data_store"
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
	addr := ":8001"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to initializa TCP listen: %v", err)
	}
	defer lis.Close()

	runGRPC(lis)
}

func runGRPC(lis net.Listener) {
	server := grpc.NewServer()

	dbConfig, err := DBConfigFromViper()
	if err != nil {
		panic(err.Error())
	}
	db := data_store.New(dbConfig)
	pb.RegisterAccountsServer(server, srv.NewServer(db))

	log.Printf("gRPC Listening on %s\n", lis.Addr().String())
	server.Serve(lis)
}
