package httpserver

import (
	"fmt"
	"net/http"

	interactions "go-specs-greet/v2/domain/interactions"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, interactions.Greet(name))
}
