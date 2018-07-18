package promote

import (
	pb "github.com/chuckleheads/hurtlocker/components/logsrv/logrecv"
	"github.com/chuckleheads/hurtlocker/components/tasker/pkg/cmd"
	"github.com/spf13/viper"
)

type PromoteCli struct {
	logsrv pb.LogRecv_ReceiveLogsClient
}

func New(logsrv pb.LogRecv_ReceiveLogsClient) PromoteCli {
	return PromoteCli{
		logsrv,
	}
}

func (b *PromoteCli) Promote() {
	cmd.RunCommand(b.logsrv, "hab", "pkg", "promote", viper.GetString("package"), viper.GetString("channel"))
}
