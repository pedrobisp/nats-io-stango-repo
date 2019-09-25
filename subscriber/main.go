package main

import (
	"log"
	"net/http"
	"time"

	stan "github.com/nats-io/stan.go"
)

func getRequest(m *stan.Msg) {
	log.Printf("Sending request")
	res, err := http.Get("https://google.com")
	if err != nil {
		log.Fatalf("Error in sending request: %v", err)
	}
	log.Printf("Response: %v", res)
}

func main() {
	sc, err := stan.Connect("test-cluster", "subscriber", stan.NatsURL("http://localhost:4222"))
	if err != nil {
		log.Fatalf("Cannot connect to nats streaming server: %v", err)
	}
	defer sc.Close()
	_, err = sc.Subscribe("backend", getRequest, stan.StartAtTime(time.Now()))
	if err != nil {
		log.Fatalf("Error subscribing to chanel:%v", err)
	}
	for {
	}
}
