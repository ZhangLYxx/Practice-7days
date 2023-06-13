package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9998", engine))
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		//键值对
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND:%s\n", req.URL)
	}
}
