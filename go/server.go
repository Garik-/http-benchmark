package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var data string

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, data)
}

func check(c string, e error) {
	if e != nil {
		log.Fatal(c, e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("../data.json")
	check("ReadFile: ", err)
	data = string(dat)
	http.HandleFunc("/", HomeRouterHandler)
	err = http.ListenAndServe(":9000", nil)
	check("ListenAndServe: ", err)
}
