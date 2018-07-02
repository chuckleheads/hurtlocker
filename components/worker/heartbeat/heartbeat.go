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

func Start() chan<- worker.WorkerState {
	publisher, _ := zmq.NewSocket(zmq.PUB)
	if len(os.Args) == 2 {
		fmt.Println("Connecting to ", os.Args[1])
		publisher.Connect(fmt.Sprintf("tcp://%s", os.Args[1]))
	} else {
		fmt.Println("Connecting to localhost:5567")
		publisher.Connect("tcp://localhost:5567")
	}
	// create state channel
	stateTV := make(chan worker.WorkerState)

	//  Ensure subscriber connection has time to complete
	fmt.Println("Waiting for connection to complete")
	time.Sleep(time.Second)

	go pulse(publisher, stateTV)
	// send initial state
	stateTV <- worker.WorkerState_Ready
	return stateTV
}

func pulse(publisher *zmq.Socket, state chan worker.WorkerState) {
	heartbeat := &worker.Heartbeat{}
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Failed to detect hostname", err)
	}
	heartbeat.Endpoint = *proto.String(fmt.Sprintf("12345@%s", hostname))
	heartbeat.Os = worker.Os(worker.Os_value[runtime.GOOS])
	heartbeat.State = worker.WorkerState(<-state)

	// send initial pulse
	sendPulse(publisher, heartbeat)
	// continue sending pulse every 30 seconds
	tickerChan := time.NewTicker(30 * time.Second).C

	for {
		select {
		case <-tickerChan:
			sendPulse(publisher, heartbeat)
		case newState := <-state:
			heartbeat.State = worker.WorkerState(newState)
			sendPulse(publisher, heartbeat)
		}
	}
}

func sendPulse(publisher *zmq.Socket, heartbeat *worker.Heartbeat) {
	fmt.Println("Sending heartbeat state", heartbeat.GetState())
	hb, err := proto.Marshal(heartbeat)
	if err != nil {
		fmt.Println("Failed to marshall heartbeat", err)
	}
	_, err = publisher.SendBytes(hb, 0)
	if err != nil {
		fmt.Println("Failed to send message", err)
	}
}
