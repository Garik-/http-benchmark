package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
)

var data []byte

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func check(c string, e error) {
	if e != nil {
		log.Fatal(c, e)
	}
}

func main() {
	runtime.GOMAXPROCS(1) // to be more similar to node.js

	var err error
	data, err = ioutil.ReadFile("../data.json")
	check("ReadFile: ", err)
	http.HandleFunc("/", HomeRouterHandler)
	err = http.ListenAndServe(":9000", nil)
	check("ListenAndServe: ", err)
}
