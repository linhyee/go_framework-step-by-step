package main

import (
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

func handler(ctx *gee.Context) {
	ctx.String(http.StatusOK, "hell %s, you're at %s", ctx.Query("name"), ctx.Path)
}

func counter(ctx *gee.Context) {
	c++
	ctx.JSON(http.StatusOK, gee.H{"c": c})
	// ctx.HTML(http.StatusOK, "<p>hello</p>")
}
