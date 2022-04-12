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
	ctr.NatsReading(sc, channel, durableID)
	log.Fatalln(http.ListenAndServe(":8181", nil))
}
