package main

import (
	"flash"
	"fmt"
	"net/http"
)

func main() {
	server := flash.New()

	server.GET("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "Request URL: %s", req.URL.Path)
	})

	server.GET("/header", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(rw, "Header is:")
		for key, value := range req.Header {
			fmt.Fprintf(rw, "%s: %v\n", key, value)
		}
	})

	server.Run("localhost:9999")
}
