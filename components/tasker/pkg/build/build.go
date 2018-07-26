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

// Access token prefix rules:
// MUST CONTAIN AN *INVALID* base-64 character
// MUST NOT CONTAIN shell special characters (eg, !)
// SHOULD be URL-safe (just in case)
const AccessTokenPrefix = "_"

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
	writeKeyToDisk(fetchOriginKey(viper.GetString("project.origin_name")))
	os.Chdir(b.basePath)
	b.fetchDeps()
	b.cloneRepo()
	b.build()
	b.uploadArtifact()
}

func (b *BuildCli) fetchDeps() {
	exit, err := cmd.RunCommand(b.logsrv,
		"hab", "pkg", "install", "chuckleheads/fetch-code", "-b")
	catch(exit, err)
	exit, err = cmd.RunCommand(b.logsrv,
		"hab", "pkg", "install", "core/hab-backline", "core/hab-plan-build")
	catch(exit, err)
}

func (b *BuildCli) cloneRepo() {
	exit, err := cmd.RunCommand(b.logsrv,
		"hab", "pkg", "exec", "chuckleheads/fetch-code", "fetch-code",
		"--url", b.repoURL, "--path", b.basePath)
	catch(exit, err)

}

func (b *BuildCli) build() {
	exit, err := cmd.RunCommand(b.logsrv,
		"hab", "pkg", "exec", "core/hab-plan-build", "hab-plan-build",
		filepath.Join(b.basePath, filepath.Dir(b.planPath)))
	catch(exit, err)
}

func (b *BuildCli) uploadArtifact() {
	artifact, err := findHart(b.basePath)
	if err != nil {
		panic(err)
	}
	exit, err := cmd.RunCommand(b.logsrv, "hab", "pkg", "upload",
		artifact, "-c", viper.GetString("channel"))
	catch(exit, err)
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

func catch(exit int, err error) {
	if exit != 0 || err != nil {
		log.Fatalln("Failed to successfully execute command")
	}
}
