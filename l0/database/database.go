package database

import (
	"Work/model"
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "VOVQUETO3498"
	dbname   = "dbl0"
)

const (
	insertJSON   = `INSERT INTO "tablel0"("id", "name" ,"phone", "city" , "adress" , "price" , "item_name" , "brand") values($1, $2, $3 , $4 , $5 , $6 , $7 , $8)`
	selectIDJSON = `SELECT * FROM "tablel0" WHERE id=$1`
)

type Database struct {
	DbStruct *sql.DB
}

func New() Database {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatalln(err)
	}

	database := Database{
		DbStruct: db,
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return database
}

func (d Database) Add(req model.Request) {

	_, err := d.DbStruct.Exec(
		insertJSON,
		req.Id,
		req.Delivery.Name,
		req.Delivery.Phone,
		req.Delivery.City,
		req.Delivery.Address,
		req.Thing.Price,
		req.Thing.ItemName,
		req.Thing.Brand)

	if err != nil {
		log.Print(err)
	}
}

func (d Database) Get(id int) model.Request {

	row, err := d.DbStruct.Query(selectIDJSON, id)

	defer row.Close()

	p := model.Request{}

	for row.Next() {
		p = model.Request{}
		err = row.Scan(
			&p.Id,
			&p.Delivery.Name,
			&p.Delivery.Phone,
			&p.Delivery.City,
			&p.Delivery.Address,
			&p.Thing.Price,
			&p.Thing.ItemName,
			&p.Thing.Brand)
	}

	if err != nil {
		log.Print(err)
	}

	return p
}
