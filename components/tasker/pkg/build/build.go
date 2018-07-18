package build

import (
	"fmt"
	"os"
	"path/filepath"

	pb "github.com/chuckleheads/hurtlocker/components/logsrv/logrecv"
	"github.com/chuckleheads/hurtlocker/components/tasker/pkg/cmd"
)

type BuildCli struct {
	logsrv   pb.LogRecv_ReceiveLogsClient
	basePath string
	repoURL  string
	planPath string
}

func New(
	logsrv pb.LogRecv_ReceiveLogsClient,
	basePath string,
	repoURL string,
	planPath string,
) BuildCli {
	return BuildCli{
		logsrv,
		basePath,
		repoURL,
		planPath,
	}
}

func (b *BuildCli) Start() {
	os.Chdir(b.basePath)
	b.fetchDeps()
	b.cloneRepo()
	b.build()
	b.uploadArtifact()
}

func (b *BuildCli) fetchDeps() {
	fmt.Println("Fetchin ur deps")
	cmd.RunCommand(b.logsrv, "hab", "pkg", "install", "chuckleheads/fetch-code", "-b")
	cmd.RunCommand(b.logsrv, "hab", "pkg", "install", "core/hab-backline", "core/hab-plan-build")
}

func (b *BuildCli) cloneRepo() {
	fmt.Println("Clonin ur repoz")
	cmd.RunCommand(b.logsrv, "hab", "pkg", "exec", "chuckleheads/fetch-code", "fetch-code", "--url", b.repoURL, "--path", b.basePath)
}

func (b *BuildCli) build() {
	fmt.Println("buildin ur shiz")
	cmd.RunCommand(b.logsrv, "hab", "pkg", "exec", "core/hab-plan-build", "hab-plan-build", filepath.Join(b.basePath, b.planPath))
}

func (b *BuildCli) uploadArtifact() {
	fmt.Println("Uploadin dat shiz")
	artifact, err := findHart(b.basePath)
	if err != nil {
		panic(err)
	}
	cmd.RunCommand(b.logsrv, "hab", "pkg", "upload", artifact, "-c", "unstable")
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
