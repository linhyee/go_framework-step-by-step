package main

import (
	"net/http"

	"./gee"
)

func main() {
	r := gee.Default()
	r.Static("/assets", "./static")

	r.GET("/panic", func(c *gee.Context) {
		names := []string{"gee"}
		c.String(http.StatusOK, names[100])
	})
	r.GET("/index", func(c *gee.Context) {
		c.String(http.StatusOK, "this is a index page")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/hello/:name", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"name":    c.Param("name"),
				"message": c.Query("msg"),
			})
		})
	}

	r.Run(":9999")
}
