package controller

import (
	"Work/cache"
	"Work/database"
	"Work/model"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	selectAllData = `SELECT * FROM "tablel0"`
)

type Controller struct {
	database database.Database
	cache    cache.Cache
}

var isFirstCompile bool = true

func New(database database.Database, cache cache.Cache) *Controller {
	ctr := Controller{
		database: database,
		cache:    cache,
	}
	return &ctr
}

func (сtr *Controller) NatsReading(sc stan.Conn, channel string, durableID string) {

	aw, _ := time.ParseDuration("300s")
	sc.Subscribe(channel, func(msg *stan.Msg) {

		msg.Ack()

		var data model.Request

		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			log.Println("Введите json пожалуйста в поток.")
			return
		}

		log.Printf("Subscribed message from clientID for Order: %+v\n", data.Delivery.Name)
		сtr.cache.Set(data)
		сtr.database.Add(data)

	}, stan.DurableName(durableID),
		stan.MaxInflight(25),
		stan.SetManualAckMode(),
		stan.AckWait(aw),
	)
}

func (c *Controller) IndexHandler(w http.ResponseWriter, r *http.Request) {

	if isFirstCompile {
		rows, err := c.database.DbStruct.Query(selectAllData)

		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			p := model.Request{}

			err = rows.Scan(
				&p.Id,
				&p.Delivery.Name,
				&p.Delivery.Phone,
				&p.Delivery.City,
				&p.Delivery.Address,
				&p.Thing.Price,
				&p.Thing.ItemName,
				&p.Thing.Brand)

			c.cache.Set(p)
		}
		log.Println("Set in Cash Succesful")

		isFirstCompile = false
	}

	if r.Method == "POST" {

		p := model.Request{}

		idString := r.FormValue("id")

		id, _ := strconv.Atoi(idString) // приведение к int

		p, ok := c.cache.Get(id)

		log.Println(p)

		if !ok {
			p = c.database.Get(id)
			c.cache.Set(p)
		}

		products := []model.Request{}
		products = append(products, p)

		tmpl, _ := template.ParseFiles("html/index.html")
		tmpl.Execute(w, products)

	} else {
		http.ServeFile(w, r, "html/from.html")
	}

}
