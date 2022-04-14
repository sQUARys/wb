package main

import (
	"Work/cache"
	"Work/controller"
	"Work/database"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"
	"time"
)

const (
	clusterID = "test-cluster"
	clientID  = "restaurant-service"
	channel   = "test"
	durableID = "restaurant-service-durable"
)

func main() {
	db := database.New()
	c := cache.New(db)
	ctr := controller.New(db, c)
	http.HandleFunc("/", ctr.IndexHandler)
	sc, err := stan.Connect( //sc, err
		clusterID,
		clientID,
		stan.NatsURL(stan.DefaultNatsURL))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server is listening...")

	aw, _ := time.ParseDuration("300s")
	sc.Subscribe(channel, func(msg *stan.Msg) {
		ctr.NatsReading(msg)
	}, stan.DurableName(durableID),
		stan.MaxInflight(25),
		stan.SetManualAckMode(),
		stan.AckWait(aw),
	)
	log.Fatalln(http.ListenAndServe(":8181", nil))
}
