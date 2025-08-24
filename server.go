package flint

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fatih/color"
)

type Server struct {
	router   *Router
	notFound HandlerFunc
}

// Create a new flint server
func NewServer() *Server {
	s := &Server{
		router: NewRouter(),
	}
	s.notFound = func(ctx *Context) {
		ctx.Default404()
	}
	s.router.notFound = s.notFound
	return s
}

// Universal handler
func (s *Server) Handle(path string, handler HandlerFunc) {
	s.router.Handle(http.MethodGet, path, handler)
}

// GET
func (s *Server) Get(path string, handler HandlerFunc) {
	s.router.Handle(http.MethodGet, path, handler)
}

// POST
func (s *Server) Post(path string, handler HandlerFunc) {
	s.router.Handle(http.MethodPost, path, handler)
}

// PUT
func (s *Server) Put(path string, handler HandlerFunc) {
	s.router.Handle(http.MethodPut, path, handler)
}

// DELETE
func (s *Server) Delete(path string, handler HandlerFunc) {
	s.router.Handle(http.MethodDelete, path, handler)
}

// static file server
func (s *Server) Static(routePath, dir string) {
	if !strings.HasSuffix(routePath, "/") {
		routePath += "/"
	}

	s.router.Handle(http.MethodGet, routePath+"*", func(ctx *Context) {
		http.ServeFile(ctx.Writer, ctx.Request, dir+strings.TrimPrefix(ctx.Request.URL.Path, routePath))
	})
}

// NotFound override
func (s *Server) SetNotFound(handler HandlerFunc) {
	s.notFound = handler
	s.router.notFound = handler
}

// Server çalıştır
func (s *Server) Run(addr ...string) error {
	port := ":8080"
	if len(addr) > 0 && strings.TrimSpace(addr[0]) != "" {
		port = addr[0]
	}

	color.New(color.FgWhite, color.Bold).Println("Flint Server is Starting...")
	fmt.Println("---------------------------")
	color.Red("Listening on http://localhost%s", port)

	return http.ListenAndServe(port, s.router)
}
