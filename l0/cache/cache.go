package cache

import (
	"Work/database"
	"Work/model"
	"log"
	"strconv"
)

type Cache struct {
	items map[string]model.Request
	db    database.Database
}

func New(database database.Database) Cache {

	items := make(map[string]model.Request)

	// cache item
	cache := Cache{
		items: items,
		db:    database,
	}

	return cache
}

func (c *Cache) Set(value model.Request) {
	c.items[strconv.Itoa(value.Id)] = model.Request{
		Id: value.Id,
		Delivery: model.DeliveryJSON{
			Name:    value.Delivery.Name,
			City:    value.Delivery.City,
			Phone:   value.Delivery.Phone,
			Address: value.Delivery.Address,
		},
		Thing: model.ThingJSON{
			Price:    value.Thing.Price,
			ItemName: value.Thing.ItemName,
			Brand:    value.Thing.Brand,
		},
	}
}

func (c *Cache) Get(id int) (model.Request, bool) {
	item, found := c.items[strconv.Itoa(id)]
	// cache not found
	if !found {
		log.Println("Not found in Get Cash")
		return item, false
	}
	return item, true
}
