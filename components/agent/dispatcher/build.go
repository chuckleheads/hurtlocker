package dispatcher

import (
	"log"

	"github.com/chuckleheads/hurtlocker/components/agent/proto/build"
	"github.com/golang/protobuf/proto"
)

func Build(buildReq []byte) {
	build := &build.Build{}
	if err := proto.Unmarshal(buildReq, build); err != nil {
		log.Fatalln("Failed to parse Build:", err)
	}
	log.Printf("Received a message: %s", build)
}
