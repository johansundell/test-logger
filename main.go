package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		defaultHandler,
	},
}

var router *mux.Router

func main() {
	router = NewRouter()
	srv := &http.Server{
		Handler: router,
		//Addr:    httpPortInterface,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 90 * time.Second,
		ReadTimeout:  90 * time.Second,
	}

	srv.Addr = ":8080"
	log.Fatal(srv.ListenAndServe())
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "fmp-json running ;)")
}
