package main

import (
	"fmt"
	"net/http"
)

func init() {
	routes = append(routes, Route{"postHandler", "POST", "/logger/", postHandler})
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
	}
	for k, v := range r.Form {
		fmt.Println(k, v)
	}
	//fmt.Println("done")

	fmt.Fprintf(w, "Oki ;)")
}
