package main

import (
	"github.com/gorilla/mux"
	unit "unit.nginx.org/go"
)

// CreateRouter routes declared handlers to their respective paths.
func CreateRouter() *mux.Router {
	r := mux.NewRouter()

	components(r)
	r.PathPrefix("/").Handler(home)

	return r
}

func main() {
	if e := unit.ListenAndServe(":8080", CreateRouter()); e != nil {
		panic(e)
	}
}
