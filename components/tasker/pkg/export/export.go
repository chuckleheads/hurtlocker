package export

import (
	"context"
	"log"
	"time"

	pb "github.com/chuckleheads/hurtlocker/components/logsrv/logrecv"
	"github.com/chuckleheads/hurtlocker/components/tasker/pkg/cmd"
	"github.com/spf13/viper"
)

type ExportCli struct {
	logsrv pb.LogRecv_ReceiveLogsClient
	ident  string
}

func New() ExportCli {
	exportCli := ExportCli{
		ident: viper.GetString("ident"),
	}
	if viper.GetBool("enable_log_stream") {
		client := cmd.LogSrvClient()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		stream, err := client.ReceiveLogs(ctx)
		if err != nil {
			log.Fatalf("%v.RecordRoute(_) = _, %v", client, err)
		}
		exportCli.logsrv = stream
	}
	return exportCli
}

func (e *ExportCli) Export() {
	cmd.RunCommand(e.logsrv, "hab", "pkg", "export", "docker", e.ident)
}
