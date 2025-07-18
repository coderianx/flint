package flint

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"sync"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (c *Context) FormData(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) FormInt(key string) int {
	val := c.FormData(key)
	num, _ := strconv.Atoi(val)
	return num
}

func (c *Context) FormMD5(key string) string {
	hash := md5.Sum([]byte(c.FormData(key)))
	return hex.EncodeToString(hash[:])
}

func (c *Context) FormSHA256(key string) string {
	hash := sha256.Sum256([]byte(c.FormData(key)))
	return hex.EncodeToString(hash[:])
}

func (c *Context) FormSHA512(key string) string {
	hash := sha512.Sum512([]byte(c.FormData(key)))
	return hex.EncodeToString(hash[:])
}

func (c *Context) Post() bool {
	return c.Request.Method == http.MethodPost
}

func (c *Context) Get() bool {
	return c.Request.Method == http.MethodGet
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

// Template cache
var tmplCache = make(map[string]*template.Template)
var tmplLock sync.RWMutex

func (c *Context) HTML(status int, tmplPath string, data interface{}) {
	tmplLock.RLock()
	tmpl := tmplCache[tmplPath]
	tmplLock.RUnlock()

	if tmpl == nil {
		t, err := template.ParseFiles(tmplPath)
		if err != nil {
			http.Error(c.Writer, err.Error(), 500)
			return
		}
		tmplLock.Lock()
		tmplCache[tmplPath] = t
		tmplLock.Unlock()
		tmpl = t
	}

	c.Writer.Header().Set("Content-Type", "text/html")
	c.Writer.WriteHeader(status)
	tmpl.Execute(c.Writer, data)
}
