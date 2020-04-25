package flash

import (
	"fmt"
	"net/http"
)

// HandlerFunc defines handler function
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine defines the main struct of http server
type Engine struct {
	router map[string]HandlerFunc
}

func (engine *Engine) addRoute(key string, handler HandlerFunc) {
	engine.router[key] = handler
}

// GET adds Get method handler function to http server
func (engine *Engine) GET(url string, handler HandlerFunc) {
	key := "GET" + url
	engine.addRoute(key, handler)
}

// POST adds Post method handler function to http server
func (engine *Engine) POST(url string, handler HandlerFunc) {
	key := "POST" + url
	engine.addRoute(key, handler)
}

// New initializes http server
func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

// Run runs http server
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	key := req.Method + req.URL.Path

	if handler, ok := engine.router[key]; ok {
		handler(rw, req)
	} else {
		fmt.Fprintf(rw, "404 NOT FOUND")
	}
}
