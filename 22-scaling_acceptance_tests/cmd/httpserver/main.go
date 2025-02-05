package main

import (
	"log"
	"net/http"

	gsg "go-specs-greet/v2"
)

func main() {
	handler := http.HandlerFunc(gsg.Handler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
