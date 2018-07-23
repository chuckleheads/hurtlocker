package dispatcher

import (
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/chuckleheads/hurtlocker/components/agent/drivers"
	"github.com/chuckleheads/hurtlocker/components/agent/proto/export"
	"github.com/golang/protobuf/proto"
)

func Export(exportReq []byte) {
	export := &export.Export{}
	if err := proto.Unmarshal(exportReq, export); err != nil {
		log.Fatalln("Failed to parse Export:", err)
	}

	config, err := json.Marshal(&export)
	if err != nil {
		log.Fatalln(err)
	}

	runExport(base64.StdEncoding.EncodeToString(config))
}

func runExport(config string) {
	driver := drivers.New()
	driver.Pull()
	mounts := map[string]string{
		"/var/run/docker.sock": "/var/run/docker.sock",
	}
	driver.Start(driver.Create(mounts, []string{"tasker", "export", "--config-string", config}))
}
