package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	pba "github.com/chuckleheads/hurtlocker/components/originsrv/api/origins"
	"github.com/chuckleheads/hurtlocker/components/originsrv/config"
	"github.com/chuckleheads/hurtlocker/components/originsrv/data_store"
	srv "github.com/chuckleheads/hurtlocker/components/originsrv/origins/server"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
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
	pba.RegisterOriginsServer(server, srv.NewServer(db))
	log.Printf("gRPC Listening on %s\n", lis.Addr().String())
	server.Serve(lis)
}

func runHTTP(config *config.Config, clientAddr string) {
	addr := fmt.Sprintf(":%d", config.HttpPort)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	mux := runtime.NewServeMux()

	if err := pba.RegisterOriginsHandlerFromEndpoint(context.Background(), mux, clientAddr, opts); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}

	log.Printf("HTTP Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
