package main

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func main() {
	var servers string
	servers = "nats://127.0.0.1:4222,nats://127.0.0.1:14222,nats://127.0.0.1:24222"
	nc, err := nats.Connect(servers)
	if err != nil {
		log.Fatal(err)
	}
	sc, err := stan.Connect("cluster1", "123", stan.NatsConn(nc))
	if err != nil {
		log.Fatal(err)
	}
	sc.Subscribe("hello", func(m *stan.Msg) {
		log.Printf("[Received] %+v", m)
	})
	sc.Publish("hello", []byte("one"))

	select {}
}
