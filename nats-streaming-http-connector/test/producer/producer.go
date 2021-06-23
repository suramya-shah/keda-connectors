package main

import (
	"log"
	"fmt"
	"strconv"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func main() {
	nc, err := nats.Connect("nats://defaultFissionAuthToken@nats-streaming.fission:4222")
	if err != nil {
		log.Fatal(err)
	}
	sc, err := stan.Connect("fissionMQTrigger", "clientPub", stan.NatsConn(nc))
	if err != nil {
		log.Fatal(err)
	}
	for i := 100; i < 110; i++ {
		sc.Publish("request", []byte("Test"+strconv.Itoa(i)))
	}
	fmt.Println("Published all the messages")
	select {}
}
