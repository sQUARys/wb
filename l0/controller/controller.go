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
	"time"
)

const (
	selectAllData = `SELECT * FROM "datasl0"`
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

	aw, _ := time.ParseDuration("60s")
	sc.Subscribe(channel, func(msg *stan.Msg) {

		msg.Ack()

		var data model.Product

		err := json.Unmarshal(msg.Data, &data)

		if err != nil {
			log.Print(err)
			return
		}

		log.Printf("Subscribed message from clientID for Order: %+v\n", *data.FirstName)
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
			p := model.Product{}
			err = rows.Scan(&p.Id, &p.FirstName, &p.LastName)
			c.cache.Set(p)
		}
		isFirstCompile = false
	}

	if r.Method == "POST" {

		//c.cache.OnAddFromNats()

		p := model.Product{}

		id := r.FormValue("id")

		f, l, ok := c.cache.Get(id)

		if ok {
			p.FirstName = &f
			p.LastName = &l
			p.Id = &id
		} else {
			p = c.database.Get(id)
			c.database.Add(p)
		}

		products := []model.Product{}
		products = append(products, p)

		tmpl, _ := template.ParseFiles("html/index.html")
		tmpl.Execute(w, products)

	} else {
		http.ServeFile(w, r, "html/from.html")
	}

}
