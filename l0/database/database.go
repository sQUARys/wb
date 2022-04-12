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
	insertToDB = `INSERT INTO "datasl0"("id","first_name", "last_name") values($1, $2, $3)`
	selectID   = `SELECT * FROM "datasl0" WHERE id=$1`
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

func (d Database) Get(id string) model.Product {

	row, err := d.DbStruct.Query(selectID, id)

	defer row.Close()

	p := model.Product{}

	for row.Next() {
		p = model.Product{}
		err = row.Scan(&p.Id, &p.FirstName, &p.LastName)
	}
	if err != nil {
		log.Print(err)
	}

	return p
}

func (d Database) Add(product model.Product) {
	_, err := d.DbStruct.Exec(insertToDB, *product.Id, *product.FirstName, *product.LastName)
	if err != nil {
		log.Print(err)
	}
}
