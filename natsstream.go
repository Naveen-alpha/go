package main

import (
	"encoding/json"
	"fmt"

	nats "github.com/nats-io/go-nats"
	stan "github.com/nats-io/go-nats-streaming"
	//stan "github.com/nats-io/stan.go"
)

type Book struct {
	ID   string
	Name string
}

func main() {
	// nc, err := nats.Connect(":4222")
	// fmt.Println(err)
	nc, err := nats.Connect(":4222")
	fmt.Println(err)

	sc, er := stan.Connect("test-cluster", "test", stan.NatsConn(nc))
	fmt.Println(er)
	b := Book{ID: "5", Name: "naveen"}
	bt, _ := json.Marshal(b)
	// stan.ConnectWait(10000)

	// i := stan.DefaultPingInterval
	// fmt.Println(i)
	// url := stan.DefaultOptions.NatsURL
	// fmt.Println(url)
	// sc, er := stan.Connect(":4222", ":4222")
	// fmt.Println(er)
	// Simple Async Subscriber
	// sc.Subscribe("foo", func(m *stan.Msg) {
	// 	fmt.Printf("Received a message: %s\n", string(m.Data))
	// })

	// Simple Synchronous Publisher
	sc.Publish("hii", bt)

	// does not return until an ack has been received from NATS Streaming

	// Subscribe starting with most recently published value
	// sc.Subscribe("foo", func(m *stan.Msg) {
	// 	fmt.Printf("Received a message: %s\n", string(m.Data))
	// }, stan.StartWithLastReceived())

	// Receive all stored values in order
	sc.Subscribe("hii", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	}, stan.DeliverAllAvailable())
	// sc.Subscribe("foo", func(m *stan.Msg) {
	// 	fmt.Printf("Received a message: %s\n", string(m.Data))
	// }, stan.DeliverAllAvailable())

	// // Receive messages starting at a specific sequence number
	// sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
	// 	fmt.Printf("Received a message: %s\n", string(m.Data))
	// }, stan.StartAtSequence(22))

	// // Subscribe starting at a specific time
	// var startTime time.Time	select {}

	// sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
	// 	fmt.Printf("Received a message: %s\n", string(m.Data))
	// }, stan.StartAtTime(startTime))

	// // Subscribe starting a specific amount of time in the past (e.g. 30 seconds ago)
	// sub, err := sc.Subscribe("foo", func(m *stan.Msg) {	select {}
	// 	fmt.Printf("Received a message: %s\n", string(m.Data))
	// }, stan.StartAtTimeDelta(time.ParseDuration("30s")))
	select {}
}
