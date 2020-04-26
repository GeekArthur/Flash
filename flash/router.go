package flash

import (
	"fmt"
)

// HandlerFunc defines handler function
type HandlerFunc func(c *Context)

type router struct {
	router map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		router: make(map[string]HandlerFunc),
	}
}

func (r *router) addRoute(method string, url string, handler HandlerFunc) {
	key := method + url
	r.router[key] = handler
}

func (r *router) handleRoute(c *Context) {
	key := c.Method + c.URL

	if handler, ok := r.router[key]; ok {
		handler(c)
	} else {
		fmt.Fprintf(c.Rw, "404 NOT FOUND")
	}
}
