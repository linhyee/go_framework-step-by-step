package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// origin object
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path, Method string
	Params       map[string]string
	// response info
	StatusCode int
}

// newContext new a context for http interface
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// Param get param vlue
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// PostForm get value from form
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query get url param's value
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader
func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

// String string for plain text output
func (c *Context) String(code int, format string, values ...interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/plain")
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON json ouput
func (c *Context) JSON(code int, obj interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data bytes output
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML html output
func (c *Context) HTML(code int, html string) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/html")
	c.Writer.Write([]byte(html))
}
