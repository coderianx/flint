package flint

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

type Server struct {
	router   *Router
	notFound HandlerFunc
}

// A new Flint Server is created.
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

	boldWhite := color.New(color.FgWhite, color.Bold)
	boldWhite.Println("Flint Server is Starting...")
	fmt.Println("---------------------------")

	displayAddr := port
	if strings.HasPrefix(port, ":") {
		displayAddr = "http://localhost" + port
	} else if strings.HasPrefix(port, "0.0.0.0") {
		displayAddr = "http://" + strings.Replace(port, "0.0.0.0", "localhost", 1)
	} else {
		displayAddr = "http://" + port
	}

	color.Red("Listening on %s", displayAddr)
	color.Cyan("Port: %s", strings.TrimPrefix(port, ":"))

	err := http.ListenAndServe(port, s.router)
	if err != nil {
		LogError("Run()", err.Error())
	}
	return err
}
