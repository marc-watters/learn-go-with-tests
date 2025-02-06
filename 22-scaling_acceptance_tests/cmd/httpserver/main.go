package main

import (
	"go-specs-greet/v2/adapters/httpserver"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", httpserver.NewHandler()))
}
