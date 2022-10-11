package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func FromRoute(name string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[name]
}
