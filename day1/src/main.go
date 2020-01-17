package main

import (
	"fmt"
	"net/http"

	"./gee"
)

var (
	c = 0
)

func main() {
	r := gee.New()
	r.GET("/", gee.HandlerFunc(handler))
	r.GET("/count", gee.HandlerFunc(counter))

	r.Run(":9999")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	c++
	fmt.Fprintf(w, "n=%d\n", c)
}
