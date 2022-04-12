package controller

import (
	"Work/cache"
	"Work/database"
	"Work/model"
	"html/template"
	"log"
	"net/http"
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

		c.cache.OnAddFromNats()

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
