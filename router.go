// router.go
package flint

import (
	"net/http"
	"strings"
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
	methodHandlers, ok := r.routes[req.Method]
	if !ok {
		if r.notFound != nil {
			ctx := &Context{Writer: w, Request: req}
			r.notFound(ctx)
			return
		}
		http.NotFound(w, req)
		return
	}

	// statik veya parametreli kontrol
	for path, handler := range methodHandlers {
		params := make(map[string]string)
		if matchPath(path, req.URL.Path, params) {
			ctx := &Context{Writer: w, Request: req, Params: params}
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

// path ile request path eşleşiyor mu, parametreleri doldur
func matchPath(routePath, reqPath string, params map[string]string) bool {
	routeParts := splitPath(routePath)
	reqParts := splitPath(reqPath)

	if len(routeParts) != len(reqParts) {
		return false
	}

	for i := range routeParts {
		if len(routeParts[i]) > 0 && routeParts[i][0] == ':' {
			// parametreyi al
			params[routeParts[i][1:]] = reqParts[i]
		} else if routeParts[i] != reqParts[i] {
			return false
		}
	}

	return true
}

func splitPath(path string) []string {
	path = strings.Trim(path, "/")
	if path == "" {
		return []string{}
	}
	return strings.Split(path, "/")
}
