package httpserver

import (
	"fmt"
	"net/http"

	interactions "go-specs-greet/v2/domain/interactions"
)

const (
	greetPath     = "/greet"
	introducePath = "/introduce"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(greetPath, replyWith(interactions.Greet))
	mux.HandleFunc(introducePath, replyWith(interactions.Introduce))
	return mux
}

func replyWith(f func(name string) (interaction string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Fprint(w, f(name))
	}
}
