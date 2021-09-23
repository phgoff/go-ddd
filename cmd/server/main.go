package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/phgoff/go-ddd/pkg/adding"
	"github.com/phgoff/go-ddd/pkg/http/rest"
	"github.com/phgoff/go-ddd/pkg/listing"
	"github.com/phgoff/go-ddd/pkg/reviewing"
	"github.com/phgoff/go-ddd/pkg/storage/json"
)

func main() {

	var lister listing.Service
	var adder adding.Service
	var reviewer reviewing.Service

	s, err := json.NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	lister = listing.NewService(s)
	adder = adding.NewService(s)
	reviewer = reviewing.NewService(s)

	router := rest.Handler(lister, adder, reviewer)

	fmt.Println("Starting server on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
