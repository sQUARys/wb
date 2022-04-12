package cache

import (
	"Work/database"
	"Work/model"
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
