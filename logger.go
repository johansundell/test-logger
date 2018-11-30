package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func init() {
	routes = append(routes, Route{"postHandler", "POST", "/logger/", postHandler})
}

type FmsData struct {
	Filename string `json:"filename"`
	Layout   string `json:"layout"`
	User     string `json:"user"`
	Action   string `json:"action"`
	Recid    string `json:"id"`
	Fields   string `json:"fields"`
}

func (data FmsData) toString() string {
	return fmt.Sprintf("%+v", data)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
	}
	if formData, found := r.Form["data"]; found {
		data := FmsData{}
		fmt.Println(formData[0])
		if err := json.Unmarshal([]byte(formData[0]), &data); err != nil {
			fmt.Println("Got data, but wrong format")
			return
		}
		//fmt.Println(data.toString())
		log.Println(data.toString())
	}

	fmt.Fprintf(w, "Oki ;)")
}
