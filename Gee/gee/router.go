package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlders map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlders: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlders[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlders[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusOK, "404 NOT FOUND.\n")
	}
}
