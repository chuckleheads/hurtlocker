package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/chuckleheads/hurtlocker/components/logsrv/logrecv"
	pbr "github.com/chuckleheads/hurtlocker/components/logsrv/logrecv/request"
	"github.com/go-cmd/cmd"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var path string

// startCmd represents the start command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Execute a Habitat build",
	Run: func(cobraCmd *cobra.Command, args []string) {
		client := logsrvClient()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		stream, err := client.ReceiveLogs(ctx)
		if err != nil {
			log.Fatalf("%v.RecordRoute(_) = _, %v", client, err)
		}
		// Disable output buffering, enable streaming
		cmdOptions := cmd.Options{
			Buffered:  false,
			Streaming: true,
		}

		// TED TODO: Set a HAB_ORIGIN

		// Create Cmd with options
		habCmd := cmd.NewCmdOptions(cmdOptions, "hab", "studio", "build", path)

		// Print STDOUT and STDERR lines streaming from Cmd
		go func() {
			for {
				select {
				case line := <-habCmd.Stdout:
					fmt.Println(line)

					stream.Send(&pbr.LogLine{
						StdoutLine: line,
					})
				case line := <-habCmd.Stderr:
					fmt.Fprintln(os.Stderr, line)
					stream.Send(&pbr.LogLine{
						StderrLine: line,
					})
				}
			}
		}()

		// Run and wait for Cmd to return, discard Status
		<-habCmd.Start()

		// Cmd has finished but wait for goroutine to print all lines
		for len(habCmd.Stdout) > 0 || len(habCmd.Stderr) > 0 {
			time.Sleep(10 * time.Millisecond)
			reply, err := stream.CloseAndRecv()
			if err != nil {
				log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
			}
			log.Printf("Route summary: %v", reply)
		}
	},
}

func init() {
	buildCmd.Flags().StringVar(&path, "path", ".", "build path (default is: '.')")
	rootCmd.AddCommand(buildCmd)
}

func logsrvClient() pb.LogRecvClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(":9211", opts...)
	if err != nil {
		panic(err.Error())
	}
	return pb.NewLogRecvClient(conn)
}
