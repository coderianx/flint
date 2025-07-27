package flint

import (
	"fmt"
	"net/http"
)

type Server struct {
	router   *Router
	notFound HandlerFunc
}

func NewServer() *Server {
	s := &Server{
		router: NewRouter(),
	}
	// Varsayılan 404
	s.notFound = func(ctx *Context) {
		ctx.Default404()
	}
	s.router.notFound = s.notFound
	return s
}

func (s *Server) Handle(path string, handler HandlerFunc) {
	s.router.Handle(path, handler)
}

func (s *Server) Static(urlPath string, dir string) {
	fs := http.FileServer(http.Dir(dir))
	s.router.Handle(urlPath, func(ctx *Context) {
		http.StripPrefix(urlPath, fs).ServeHTTP(ctx.Writer, ctx.Request)
	})
}

func (s *Server) SetNotFound(handler HandlerFunc) {
	s.notFound = handler
	s.router.notFound = handler
}

/*
func (s *Server) Run(addr string) error {
	fmt.Println("Flint - v1.0\nServer is Starting...")
	return http.ListenAndServe(addr, s.router)
}
*/

func (s *Server) Run(addr string) error {
	fmt.Println("Flint-Server starting...")
	fmt.Println(addr)
	err := http.ListenAndServe(addr, s.router)

	if err != nil {
		LogError("Run()", err.Error())
	}
	return err
}
