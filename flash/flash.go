package flash

import (
	"net/http"
)

// Engine defines the main struct of http server
type Engine struct {
	router *router
}

// GET adds Get method handler function to http server
func (engine *Engine) GET(url string, handler HandlerFunc) {
	engine.router.addRoute("GET", url, handler)
}

// POST adds Post method handler function to http server
func (engine *Engine) POST(url string, handler HandlerFunc) {
	engine.router.addRoute("POST", url, handler)
}

// New initializes http server
func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

// Run runs http server
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	context := NewContext(rw, req)
	engine.router.handleRoute(context)
}
