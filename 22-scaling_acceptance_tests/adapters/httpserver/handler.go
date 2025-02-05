package httpserver

import (
	"fmt"
	"net/http"

	gsg "go-specs-greet/v2"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, gsg.Greet(name))
}
