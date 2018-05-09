package cmd

import (
	"log"
	"net"

	"github.com/chuckleheads/hurtlocker/components/originsrv/data_store"
	pb "github.com/chuckleheads/hurtlocker/components/originsrv/origins"
	srv "github.com/chuckleheads/hurtlocker/components/originsrv/origins/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		setup()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func setup() {
	addr := ":7001"
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
	pb.RegisterOriginsServer(server, srv.NewServer(db))

	log.Printf("gRPC Listening on %s\n", lis.Addr().String())
	server.Serve(lis)
}
