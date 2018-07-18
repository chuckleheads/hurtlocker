package promote

import (
	"context"
	"log"
	"time"

	pb "github.com/chuckleheads/hurtlocker/components/logsrv/logrecv"
	"github.com/chuckleheads/hurtlocker/components/tasker/pkg/cmd"
	"github.com/spf13/viper"
)

type PromoteCli struct {
	logsrv pb.LogRecv_ReceiveLogsClient
}

func New() PromoteCli {
	promoteCli := PromoteCli{}
	if viper.GetBool("enable_log_stream") {
		client := cmd.LogSrvClient()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		stream, err := client.ReceiveLogs(ctx)
		if err != nil {
			log.Fatalf("%v.RecordRoute(_) = _, %v", client, err)
		}
		promoteCli.logsrv = stream
	}
	return promoteCli
}

func (b *PromoteCli) Promote() {
	cmd.RunCommand(b.logsrv, "hab", "pkg", "promote", viper.GetString("package"), viper.GetString("channel"))
}
