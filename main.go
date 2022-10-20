package main

import (
	"log"
	"net/http"
	"valio/routes"
)

func main() {

	

	if err := http.ListenAndServe(":8080", routes.Routes()); err != nil {
        log.Fatal(err)
    }

}
