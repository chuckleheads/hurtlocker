package dispatcher

import (
	"log"

	"github.com/chuckleheads/hurtlocker/components/agent/proto/build"
	"github.com/go-cmd/cmd"
	"github.com/golang/protobuf/proto"
)

func Build(buildReq []byte) {
	build := &build.Build{}
	if err := proto.Unmarshal(buildReq, build); err != nil {
		log.Fatalln("Failed to parse Build:", err)
	}
	runBuild(build)
}

func runBuild(payload *build.Build) {
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: false,
	}
	tasker := cmd.NewCmdOptions(cmdOptions, "hab", "pkg", "exec", "chuckleheads/tasker", "tasker", "build", payload.GetPackagePath())
	// Fire and forget - we don't care about the status because we are just dispatching a task runner
	tasker.Start()
}
