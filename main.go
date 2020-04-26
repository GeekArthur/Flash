package main

import (
	"flash"
	"net/http"
)

func main() {
	server := flash.New()

	server.GET("/html", func(c *flash.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	server.GET("/string", func(c *flash.Context) {
		testString := "Hello World"
		c.String(http.StatusOK, "%s", testString)
	})

	server.GET("/json", func(c *flash.Context) {
		c.JSON(http.StatusOK,
			`{
			Name: "Arthur",
			Age: "20",
			Occupation: "Softare Engineer",
		}`)
	})

	server.Run("localhost:9999")
}
