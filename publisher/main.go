package main

import (
	"log"

	stan "github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect(
		"test-cluster",
		"publisher",
		stan.NatsURL("http://localhost:4222"),
	)
	if err != nil {
		log.Fatalf("Error connecting to streaming server:%v", err)
	}
	defer sc.Close()
	err = sc.Publish("backend", []byte("Hello"))
	if err != nil {
		log.Fatalf("Error publishing to backend channel")
	}
}
