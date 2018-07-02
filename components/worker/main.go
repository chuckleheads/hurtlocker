package main

import (
	"github.com/chuckleheads/hurtlocker/components/worker/heartbeat"
)

func main() {
	forever := make(chan bool)
	heartbeat.Start()
	// Uncomment for demo\
	// pulse := heartbeat.Start()
	// time.Sleep(5 * time.Second)
	// pulse <- worker.WorkerState_Busy
	// time.Sleep(5 * time.Second)
	// pulse <- worker.WorkerState_Ready
	<-forever
}
