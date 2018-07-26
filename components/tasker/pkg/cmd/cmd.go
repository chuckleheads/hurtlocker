package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/chuckleheads/hurtlocker/components/logsrv/logrecv"
	pbr "github.com/chuckleheads/hurtlocker/components/logsrv/logrecv/request"

	"github.com/go-cmd/cmd"
	"github.com/spf13/viper"
)

func RunCommand(stream pb.LogRecv_ReceiveLogsClient, command ...string) (int, error) {
	// Disable output buffering, enable streaming
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}

	// Create Cmd with options
	kommand, args := command[0], command[1:]
	habCmd := cmd.NewCmdOptions(cmdOptions, kommand, args...)
	//filthy hack because I'm lazy
	origin := fmt.Sprintf("HAB_ORIGIN=%s", viper.GetString("project.origin_name"))
	bldr := fmt.Sprintf("HAB_BLDR_URL=%s", viper.GetString("bldr-url"))
	habCmd.Env = []string{origin, bldr}
	// Print STDOUT and STDERR lines streaming from Cmd
	go func() {
		for {
			select {
			case line := <-habCmd.Stdout:
				fmt.Println(line)

				sendStream(stream, pbr.LogLine{
					StdoutLine: line,
				})
			case line := <-habCmd.Stderr:
				fmt.Fprintln(os.Stderr, line)
				sendStream(stream, pbr.LogLine{
					StderrLine: line,
				})
			}
		}
	}()

	// Run and wait for Cmd to return, discard Status
	<-habCmd.Start()

	// Cmd has finished but wait for goroutine to print all lines
	if stream != nil {
		for len(habCmd.Stdout) > 0 || len(habCmd.Stderr) > 0 {
			time.Sleep(10 * time.Millisecond)
			reply, err := stream.CloseAndRecv()
			if err != nil {
				log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
			}
			log.Printf("Route summary: %v", reply)
		}
	}
	return habCmd.Status().Exit, habCmd.Status().Error
}

func sendStream(stream pb.LogRecv_ReceiveLogsClient, line pbr.LogLine) {
	if stream != nil {
		stream.Send(&line)
	}
}
