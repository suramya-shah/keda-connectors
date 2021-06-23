package main

import (
	"log"
	"strconv"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func main() {
	nc, err := nats.Connect("nats://defaultFissionAuthToken@nats-streaming:4222")
	if err != nil {
		log.Fatal(err)
	}
	sc, err := stan.Connect("fissionMQTrigger", "clientPub", stan.NatsConn(nc))
	if err != nil {
		log.Fatal(err)
	}
	for i := 100; i < 200; i++ {
		sc.Publish("request", []byte("Test"+strconv.Itoa(i)))
	}

	select {}
}
