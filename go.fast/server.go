package main

import (
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"log"
	"runtime"
)

func check(c string, e error) {
	if e != nil {
		log.Fatal(c, e)
	}
}

func main() {
	runtime.GOMAXPROCS(1) // to be more similar to node.js

	data, err := ioutil.ReadFile("../data.json")
	check("ReadFile: ", err)

	err = fasthttp.ListenAndServe(":9000", func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("application/json")
		ctx.Response.SwapBody(data)
	})
	check("ListenAndServe: ", err)
}
