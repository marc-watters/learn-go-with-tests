package main

import (
	"go-specs-greet/v2/adapters/httpserver"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(httpserver.Handler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
