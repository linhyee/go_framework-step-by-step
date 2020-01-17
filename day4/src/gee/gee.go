package gee

import (
	"log"
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(*Context)

// Engine struct implement the interface of ServeHTTP
type (
	Engine struct {
		*RouterGroup // abstract engine as group
		router       *router
		groups       []*RouterGroup // store all groups
	}

	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc // support middleware
		parent      *RouterGroup  // support nesting
		engine      *Engine       // all groups share a Engine instance
	}
)

// New function is the constructor of gee.Engine
func New() *Engine {
	e := &Engine{router: newRouter()}
	e.RouterGroup = &RouterGroup{engine: e}
	e.groups = []*RouterGroup{e.RouterGroup}
	return e
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (g *RouterGroup) Group(prefix string) *RouterGroup {
	e := g.engine
	ng := &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: e,
	}
	e.groups = append(e.groups, ng)
	return ng
}

// addRoute add a new route
func (g *RouterGroup) addRoute(method, comp string, handler HandlerFunc) {
	pattern := g.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	g.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (g *RouterGroup) GET(pattern string, handler HandlerFunc) {
	g.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (g *RouterGroup) POST(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}

// Run defiens the method to start a http server
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}
