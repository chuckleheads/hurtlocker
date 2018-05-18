package dispatcher

import (
	"log"

	"github.com/chuckleheads/hurtlocker/components/agent/proto/deploy"
	"github.com/golang/protobuf/proto"
)

func Deploy(deployReq []byte) {
	deploy := &deploy.Deploy{}
	if err := proto.Unmarshal(deployReq, deploy); err != nil {
		log.Fatalln("Failed to parse Deploy:", err)
	}
	log.Printf("Received a message: %s", deploy)
}
