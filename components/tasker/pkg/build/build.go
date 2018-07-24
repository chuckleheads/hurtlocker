package build

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	pb "github.com/chuckleheads/hurtlocker/components/logsrv/logrecv"
	"github.com/chuckleheads/hurtlocker/components/tasker/pkg/cmd"
	"github.com/spf13/viper"
)

type BuildCli struct {
	logsrv   pb.LogRecv_ReceiveLogsClient
	basePath string
	repoURL  string
	planPath string
}

func New(basePath string) BuildCli {
	buildCli := BuildCli{
		basePath: basePath,
		repoURL:  viper.GetString("project.vcs_data"),
		planPath: viper.GetString("project.plan_path"),
	}
	if viper.GetBool("enable_log_stream") {
		client := cmd.LogSrvClient()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		stream, err := client.ReceiveLogs(ctx)
		if err != nil {
			log.Fatalf("%v.RecordRoute(_) = _, %v", client, err)
		}
		buildCli.logsrv = stream
	}
	return buildCli
}

func (b *BuildCli) Start() {
	// TED TODO: should be checking for errors here and bailing early
	os.Chdir(b.basePath)
	b.fetchDeps()
	b.cloneRepo()
	b.build()
	b.uploadArtifact()
}

func (b *BuildCli) fetchDeps() {
	cmd.RunCommand(b.logsrv, "hab", "pkg", "install", "chuckleheads/fetch-code", "-b")
	cmd.RunCommand(b.logsrv, "hab", "pkg", "install", "core/hab-backline", "core/hab-plan-build")
}

func (b *BuildCli) cloneRepo() {
	cmd.RunCommand(b.logsrv, "hab", "pkg", "exec", "chuckleheads/fetch-code", "fetch-code", "--url", b.repoURL, "--path", b.basePath)
}

func (b *BuildCli) build() {
	cmd.RunCommand(b.logsrv, "hab", "pkg", "exec", "core/hab-plan-build", "hab-plan-build", filepath.Join(b.basePath, filepath.Dir(b.planPath)))
}

func (b *BuildCli) uploadArtifact() {
	artifact, err := findHart(b.basePath)
	if err != nil {
		panic(err)
	}
	cmd.RunCommand(b.logsrv, "hab", "pkg", "upload", artifact, "-c", viper.GetString("channel"))
}

func findHart(path string) (string, error) {
	matches, err := filepath.Glob(filepath.Join(path, "results", "*.hart"))
	if err != nil {
		return "", err
	}
	if len(matches) == 0 {
		return "", fmt.Errorf("No Matches Found in %s", path)
	}
	return matches[0], nil
}
