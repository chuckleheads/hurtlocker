package dispatcher

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"github.com/chuckleheads/hurtlocker/components/agent/drivers"
	"github.com/chuckleheads/hurtlocker/components/agent/proto/build"

	"github.com/golang/protobuf/proto"
)

func Build(buildReq []byte) {
	build := &build.BuildRequest{}
	if err := proto.Unmarshal(buildReq, build); err != nil {
		log.Fatalln("Failed to parse Build:", err)
	}

	config, err := json.Marshal(&build)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(config))

	runBuild(base64.StdEncoding.EncodeToString(config))
}

func runBuild(config string) {
	driver := drivers.New()
	driver.Pull()
	// TED TODO - Tasker should call out to builder api to get the correct keys
	// a builder token should be provided by the agent which is generated from
	// the shared builder key
	mounts := map[string]string{
		"/hab/cache/keys": "/hab/cache/keys",
	}
	driver.Start(driver.Create(mounts, []string{"tasker", "build", "--config-string", config}))
}
