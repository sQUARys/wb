package main

import (
	"Work/model"
	"encoding/json"
	"fmt"
	stan "github.com/nats-io/stan.go"
	"log"
)

const (
	clusterID = "test-cluster"
	clientID  = "test_client"
	channel   = "test"
	durableID = "restaurant-service-durable"
)

var (
	Sc stan.Conn
)

func main() {

	data := model.Request{}
	data.Id = 5
	data.Delivery.Name = "Alex"
	data.Delivery.City = "Saint-Petersburg"
	data.Delivery.Phone = "8901234932"
	data.Delivery.Address = "ul. Karpovki"

	data.Thing.ItemName = "Socks"
	data.Thing.Price = 5
	data.Thing.Brand = "Adidas"

	var jsonData []byte
	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(jsonData))
	quit := make(chan struct{})
	ConnectStan(clientID)

	PublishNats(jsonData, channel)

	<-quit
}

func ConnectStan(clientID string) {

	// you can set client id anything
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(stan.DefaultNatsURL),
		stan.Pings(1, 3),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))

	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, stan.DefaultNatsURL)
	}

	log.Println("Connected Nats")
	Sc = sc

}

func PublishNats(data []byte, channel string) {
	ach := func(s string, err2 error) {}
	_, err := Sc.PublishAsync(channel, data, ach)
	if err != nil {
		log.Fatalf("Error during async publish: %v\n", err)
	}
}
