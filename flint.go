package flint

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (c *Context) File(filepath string) {
	http.ServeFile(c.Writer, c.Request, filepath)
}

func (c *Context) JSON(status int, data interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	json.NewEncoder(c.Writer).Encode(data)
}

func (c *Context) String(status int, text string) {
	c.Writer.Header().Set("Content-Type", "text/plain")
	c.Writer.WriteHeader(status)
	c.Writer.Write([]byte(text))
}

type HandlerFunc func(ctx *Context)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	return &Server{mux: http.NewServeMux()}
}

func (s *Server) ServerHandler(path string, handler HandlerFunc) {
	s.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			http.NotFound(w, r)
			return
		}
		ctx := &Context{Writer: w, Request: r}
		handler(ctx)
	})
}

func (s *Server) Static(urlPath string, dir string) {
	fs := http.FileServer(http.Dir(dir))
	s.mux.Handle(urlPath+"/", http.StripPrefix(urlPath, fs))
}

func (s *Server) Serve(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}

func (c *Context) HTML(status int, tmplPath string, data interface{}) {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(c.Writer, "Template parse error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	c.Writer.Header().Set("Content-Type", "text/html")
	c.Writer.WriteHeader(status)
	tmpl.Execute(c.Writer, data)
}
