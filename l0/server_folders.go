package main

import (
	"Work/cache"
	"Work/controller"
	"Work/database"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	db := database.New()
	c := cache.New(db)
	ctr := controller.New(db, c)

	http.HandleFunc("/", ctr.IndexHandler)

	fmt.Println("Server is listening...")
	log.Fatalln(http.ListenAndServe(":8181", nil))
}
