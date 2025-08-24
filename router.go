package flint

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Router struct {
	routes   map[string]map[string]HandlerFunc // method -> path -> handler
	notFound HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes:   make(map[string]map[string]HandlerFunc),
		notFound: nil,
	}
}

// Tüm methodlar için
func (r *Router) Handle(method, path string, handler HandlerFunc) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]HandlerFunc)
	}
	r.routes[method][path] = handler
}

// stdlib uyumu: ServeHTTP
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handlers, ok := r.routes[req.Method]; ok {
		if handler, ok := handlers[req.URL.Path]; ok {
			ctx := &Context{Writer: w, Request: req}
			handler(ctx)
			return
		}
	}

	if r.notFound != nil {
		ctx := &Context{Writer: w, Request: req}
		r.notFound(ctx)
		return
	}

	http.NotFound(w, req)
}
