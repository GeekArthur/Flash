package flash

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Context encapsulates http request and response
type Context struct {
	Rw     http.ResponseWriter
	Req    *http.Request
	Method string
	URL    string
	Status int
}

// NewContext initializes new context
func NewContext(rw http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Rw:     rw,
		Req:    req,
		Method: req.Method,
		URL:    req.URL.Path,
	}
}

// SetStatus sets http response status code
func (c *Context) SetStatus(status int) {
	c.Status = status
	c.Rw.WriteHeader(status)
}

// SetHeader sets http response header
func (c *Context) SetHeader(key string, value string) {
	c.Rw.Header().Set(key, value)
}

// JSON forms http response body with JSON type
func (c *Context) JSON(status int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatus(status)
	encoder := json.NewEncoder(c.Rw)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Rw, err.Error(), 500)
	}
}

// HTML forms http response body with HTML type
func (c *Context) HTML(status int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatus(status)
	c.Rw.Write([]byte(html))
}

// String forms http response body with String type
func (c *Context) String(status int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatus(status)
	c.Rw.Write([]byte(fmt.Sprintf(format, values...)))
}
