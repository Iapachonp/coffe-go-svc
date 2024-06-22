package main

import (
	"fmt"
	"net/http"
	"log"
)

const port = 8080 

type application struct {	
	Domain string

}
	
func main() {
  	// set application config 
	var app application
	// read from command line 

	// connect to db 
	app.Domain = "antesuncafe.com"


	log.Println("Starting coffe a1c app on port", port )

	// start a web server 
	err := http.ListenAndServe(fmt.Sprintf(":%d",port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
