package cmd

import (
	"log"
	"net"

	"github.com/chuckleheads/hurtlocker/components/originsrv/config"
	"github.com/chuckleheads/hurtlocker/components/originsrv/data_store"
	pb "github.com/chuckleheads/hurtlocker/components/originsrv/origins"
	srv "github.com/chuckleheads/hurtlocker/components/originsrv/origins/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// startCmd represents the start command
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

	dbConfig, err := configFromViper()
	if err != nil {
		panic(err.Error())
	}
	db := data_store.New(dbConfig)
	pb.RegisterOriginsServer(server, srv.NewServer(db))

	log.Printf("gRPC Listening on %s\n", lis.Addr().String())
	server.Serve(lis)
}

func configFromViper() (*config.DBConfig, error) {
	cfg := &config.DBConfig{}
	if err := viper.Unmarshal(cfg); err != nil {
		panic(err.Error())
	}
	return cfg, nil
}
