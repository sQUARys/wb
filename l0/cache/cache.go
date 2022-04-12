package cache

import (
	"Work/database"
	"Work/model"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

const (
	clusterID = "test-cluster"
	clientID  = "restaurant-service"
	channel   = "test"
	durableID = "restaurant-service-durable"
)

const (
	insertToDB = `INSERT INTO "datasl0"("id","first_name", "last_name") values($1, $2, $3)`
)

type Cache struct {
	items map[string]Item
	db    database.Database
}

type Item struct {
	ItemFirstName string
	ItemLastName  string
}

func New(database database.Database) Cache {

	items := make(map[string]Item)

	// cache item
	cache := Cache{
		items: items,
		db:    database,
	}

	return cache
}

func (c *Cache) Set(value model.Product) {
	c.items[*value.Id] = Item{
		ItemFirstName: *value.FirstName,
		ItemLastName:  *value.LastName,
	}
}

func (c *Cache) Get(id string) (string, string, bool) {
	item, found := c.items[id]
	// cache not found
	if !found {
		return "Not Found", "Not Found", false
	}
	return item.ItemFirstName, item.ItemLastName, true
}

func (c *Cache) OnAddFromNats() {

	sc, err := stan.Connect(
		clusterID,
		clientID,
		stan.NatsURL(stan.DefaultNatsURL))

	if err != nil {
		log.Fatal(err)
	}

	aw, _ := time.ParseDuration("60s")
	sc.Subscribe(channel, func(msg *stan.Msg) {

		msg.Ack()

		var data model.Product

		err := json.Unmarshal(msg.Data, &data)

		if err != nil {
			log.Print(err)
			return
		}

		log.Printf("Subscribed message from clientID - %s for Order: %+v\n", clientID, *data.FirstName)
		c.Set(data)
		c.db.Add(data)

	}, stan.DurableName(durableID),
		stan.MaxInflight(25),
		stan.SetManualAckMode(),
		stan.AckWait(aw),
	)
}
