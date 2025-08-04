package flint

import (
	"fmt"
	"net/http"
	"strings"
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

func (s *Server) Run(addr string) error {
	fmt.Println("Flint-Server starting...")

	displayAddr := addr
	if strings.HasPrefix(addr, ":") {
		displayAddr = "http://localhost" + addr
	} else if strings.HasPrefix(addr, "0.0.0.0") {
		displayAddr = "http://" + strings.Replace(addr, "0.0.0.0", "localhost", 1)
	} else {
		displayAddr = "http://" + addr
	}

	fmt.Println("Listening on", displayAddr)

	err := http.ListenAndServe(addr, s.router)
	if err != nil {
		LogError("Run()", err.Error())
	}
	return err
}
