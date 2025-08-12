package flint

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
)

type Server struct {
	router      *Router
	notFound    HandlerFunc
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

func (s *Server) Static(routePath, dir string) {
	if !strings.HasSuffix(routePath, "/") {
		routePath += "/"
	}

	s.router.Handle(routePath+"*", func(ctx *Context) {
		filePath := strings.TrimPrefix(ctx.Request.URL.Path, routePath)
		ctx.File(filepath.Join(dir, filePath))
	})
}

func (s *Server) SetNotFound(handler HandlerFunc) {
	s.notFound = handler
	s.router.notFound = handler
}

func (s *Server) Run(addr ...string) error {
	port := ":8080"
	if len(addr) > 0 && strings.TrimSpace(addr[0]) != "" {
		port = addr[0]
	}

	fmt.Println("Flint-Server starting...")

	displayAddr := port
	if strings.HasPrefix(port, ":") {
		displayAddr = "http://localhost" + port
	} else if strings.HasPrefix(port, "0.0.0.0") {
		displayAddr = "http://" + strings.Replace(port, "0.0.0.0", "localhost", 1)
	} else {
		displayAddr = "http://" + port
	}

	fmt.Println("Listening on", displayAddr)

	err := http.ListenAndServe(port, s.router)
	if err != nil {
		LogError("Run()", err.Error())
	}
	return err
}
