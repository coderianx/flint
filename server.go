package flint

import (
	"net/http"
)

type Server struct {
	router *Router
}

func NewServer() *Server {
	return &Server{
		router: NewRouter(),
	}
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

func (s *Server) Serve(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
