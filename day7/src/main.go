package main

import (
	"net/http"

	"./gee"
)

func main() {
	r := gee.Default()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"gee"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
