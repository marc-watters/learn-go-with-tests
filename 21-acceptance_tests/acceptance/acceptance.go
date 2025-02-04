package main

import (
	"context"
	"log"
	"net/http"

	"github.com/quii/go-graceful-shutdown/acceptancetests"

	gracefulshutdown "github.com/quii/go-graceful-shutdown"
)

func main() {
	var (
		ctx        = context.Background()
		httpServer = &http.Server{Addr: ":8080", Handler: http.HandlerFunc(acceptancetests.SlowHandler)}
		server     = gracefulshutdown.NewServer(httpServer)
	)

	if err := server.ListenAndServe(ctx); err != nil {
		// this will typcially happen if our response aren't written before the ctx deadline, not much can be done
		log.Fatalf("uh oh, didn't shutdown gracefully, some responses may have been lost %v", err)
	}

	// hopefully you'll always see this instead
	log.Println("shutdown gracefully! all responses were sent")
}
