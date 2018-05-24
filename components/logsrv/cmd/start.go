package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/chuckleheads/hurtlocker/components/logsrv/config"
	pbs "github.com/chuckleheads/hurtlocker/components/logsrv/logrecv"
	logsrv "github.com/chuckleheads/hurtlocker/components/logsrv/logrecv/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Execute a Habitat build",
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
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to initializa TCP listen: %v", err)
	}
	defer lis.Close()

	runGRPC(config, lis)
}

func runGRPC(config *config.Config, lis net.Listener) {
	server := grpc.NewServer()
	pbs.RegisterLogRecvServer(server, logsrv.NewServer())
	log.Printf("gRPC Listening on %s\n", lis.Addr().String())
	server.Serve(lis)
}
