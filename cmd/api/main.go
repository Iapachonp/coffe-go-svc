package main

import (
	"coffee-svc/internal/repository"
	"coffee-svc/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080 

type application struct {	
	DSN string
	Domain string
	DB repository.DatabaseRepo

}
	
func main() {
  	// set application config 
	var app application
	
	// read from command line
	flag.StringVar(&app.DSN, "dsn" , "host=localhost port=5432 user=postgres password=postgres dbname=coffees sslmode=disable timezone=UTC connect_timeout=5", "postgres connection string")
	flag.Parse()
	
	// connect to db
	conn , err := app.connectToDB()
	if err != nil {
		log.Fatal(err)

	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()
	
	app.Domain = "antesuncafe.com"
	log.Println("Starting coffe a1c app on port", port )

	// start a web server 
	err = http.ListenAndServe(fmt.Sprintf(":%d",port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
