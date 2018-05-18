package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/go-cmd/cmd"
	"github.com/spf13/cobra"
)

var path string

// startCmd represents the start command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Execute a Habitat build",
	Run: func(cobraCmd *cobra.Command, args []string) {
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
					// TED TODO: This should be streamed back to the API over a gRPC socket
					fmt.Println(line)
				case line := <-habCmd.Stderr:
					fmt.Fprintln(os.Stderr, line)
				}
			}
		}()

		// Run and wait for Cmd to return, discard Status
		<-habCmd.Start()

		// Cmd has finished but wait for goroutine to print all lines
		for len(habCmd.Stdout) > 0 || len(habCmd.Stderr) > 0 {
			time.Sleep(10 * time.Millisecond)
		}
	},
}

func init() {
	buildCmd.Flags().StringVar(&path, "path", ".", "build path (default is: '.')")
	rootCmd.AddCommand(buildCmd)
}
