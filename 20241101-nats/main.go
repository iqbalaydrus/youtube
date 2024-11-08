package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"time"
)

func main() {
	mode := os.Args[1]
	fmt.Println(mode)
	switch mode {
	case "pub":
		pubFunc()
	case "sub":
		subFunc()
	}
}

func pubFunc() {
	nc, err := nats.Connect("nats://192.168.88.33:4222")
	if err != nil {
		log.Panic(err)
	}
	msg, err := nc.Request("help", []byte("help me"), 10*time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received a message via request: %s\n", string(msg.Data))
}

func subFunc() {
	nc, err := nats.Connect("nats://192.168.88.33:4222")
	if err != nil {
		log.Panic(err)
	}
	nc.Subscribe("help", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
		nc.Publish(m.Reply, []byte("I can help!"))
		nc.Flush()
	})
	time.Sleep(10 * time.Second)
}

func example() {
	// Connect to a server
	nc, _ := nats.Connect("nats://192.168.88.33:4222")

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))

	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message via async: %s\n", string(m.Data))
	})

	// Responding to a request message
	nc.Subscribe("request", func(m *nats.Msg) {
		m.Respond([]byte("answer is 42"))
	})

	// Simple Sync Subscriber
	sub, err := nc.SubscribeSync("foo")
	if err != nil {
		log.Fatal(err)
	}
	m, err := sub.NextMsg(time.Second)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received a message via sync: %s\n", string(m.Data))

	// Channel Subscriber
	ch := make(chan *nats.Msg, 64)
	sub, err = nc.ChanSubscribe("foo", ch)
	if err != nil {
		log.Fatal(err)
	}
	msg := <-ch
	fmt.Printf("Received a message via channel: %s\n", string(msg.Data))

	// Unsubscribe
	sub.Unsubscribe()

	// Drain
	sub.Drain()

	// Requests
	msg, err = nc.Request("help", []byte("help me"), 10*time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received a message via request: %s\n", string(msg.Data))

	// Replies
	nc.Subscribe("help", func(m *nats.Msg) {
		nc.Publish(m.Reply, []byte("I can help!"))
	})

	// Drain connection (Preferred for responders)
	// Close() not needed if this is called.
	nc.Drain()

	// Close connection
	nc.Close()
}
