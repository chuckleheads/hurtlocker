package dispatcher

import (
	"log"

	"github.com/chuckleheads/hurtlocker/components/agent/proto/export"
	"github.com/golang/protobuf/proto"
)

func Export(exportReq []byte) {
	export := &export.Export{}
	if err := proto.Unmarshal(exportReq, export); err != nil {
		log.Fatalln("Failed to parse Export:", err)
	}
	log.Printf("Received a message: %s", export)
}
