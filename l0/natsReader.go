package main

import (
	"encoding/json"
	"log"
	"runtime"
	"time"

	stan "github.com/nats-io/stan.go"
)

type Person struct {
	Id        string
	FirstName string
	LastName  string
}

const (
	clusterID = "test-cluster"
	clientID  = "restaurant-service1"
	channel   = "test"
	durableID = "restaurant-service-durable"
)

func main() {
	sc, err := stan.Connect(
		clusterID,
		clientID,
		stan.NatsURL(stan.DefaultNatsURL))

	if err != nil {
		log.Fatal(err)
	}

	aw, _ := time.ParseDuration("60s")
	sc.Subscribe(channel, func(msg *stan.Msg) {

		msg.Ack() // Manual ACK

		var data Person

		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			log.Print(err)
			return
		}
		// Handle the message
		log.Printf("Subscribed message from clientID - %s for Order: %+v\n", clientID, data.FirstName)

	}, stan.DurableName(durableID),
		stan.MaxInflight(25),
		stan.SetManualAckMode(),
		stan.AckWait(aw),
	)
	runtime.Goexit()
}
