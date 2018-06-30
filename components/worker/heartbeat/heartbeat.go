package heartbeat

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/chuckleheads/hurtlocker/components/worker/worker"
	"github.com/golang/protobuf/proto"
	zmq "github.com/pebbe/zmq4"
)

func Start() {
	publisher, _ := zmq.NewSocket(zmq.PUB)
	if len(os.Args) == 2 {
		fmt.Println("Connecting to ", os.Args[1])
		publisher.Connect(fmt.Sprintf("tcp://%s", os.Args[1]))
	} else {
		fmt.Println("Connecting to localhost:5567")
		publisher.Connect("tcp://localhost:5567")
	}

	//  Ensure subscriber connection has time to complete
	fmt.Println("Waiting for connection to complete")
	time.Sleep(time.Second)

	heartbeat := &worker.Heartbeat{}
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Failed to detect hostname", err)
	}
	heartbeat.Endpoint = *proto.String(fmt.Sprintf("12345@%s", hostname))
	heartbeat.Os = worker.Os(worker.Os_value[runtime.GOOS])
	heartbeat.State = worker.WorkerState(worker.WorkerState_Ready)

	//  Send a heartbeat every 30 seconds
	for {
		fmt.Println("Sending heartbeat", heartbeat)
		hb, err := proto.Marshal(heartbeat)
		if err != nil {
			fmt.Println("Failed to marshall heartbeat", err)
		}
		_, err = publisher.SendBytes(hb, 0)
		if err != nil {
			fmt.Println("Failed to send message", err)
		}
		time.Sleep(30 * time.Second)
	}
}
