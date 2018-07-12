package dispatcher

import (
	"log"

	"github.com/chuckleheads/hurtlocker/components/agent/drivers"
	"github.com/chuckleheads/hurtlocker/components/agent/proto/build"

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
	driver := drivers.New()
	driver.Pull()
	driver.Start(driver.Create([]string{"build", payload.PackagePath}))
}
