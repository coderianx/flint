package flint

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Router struct {
	routes   map[string]HandlerFunc
	notFound HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes:   make(map[string]HandlerFunc),
		notFound: nil,
	}
}

func (r *Router) Handle(path string, handler HandlerFunc) {
	r.routes[path] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, ok := r.routes[req.URL.Path]; ok {
		ctx := &Context{Writer: w, Request: req}
		handler(ctx)
	} else if r.notFound != nil {
		ctx := &Context{Writer: w, Request: req}
		r.notFound(ctx)
	} else {
		http.NotFound(w, req)
	}
}
