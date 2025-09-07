/*
Package flint provides a lightweight HTTP web framework for building web applications
with Go. It offers an easy-to-use Context struct for handling requests and responses,
including JSON, HTML templates, file serving, form parsing, and HTTP method helpers.

The Context struct allows developers to access query parameters, form data, user agents,
and provides methods for sending various types of responses such as JSON, plain text,
files, and HTML templates. It also includes helpers for common HTTP methods and redirects.

Package flint aims to simplify web development by providing clear, concise, and efficient
tools while keeping the API minimal and easy to understand.
*/

package flint

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
)

// Context represents the context of a single HTTP request in the Flint framework.
// It provides convenient access to the request data, query parameters, form values,
// HTTP method checks, and response writers for sending various types of responses.
// By using Context, developers can easily read input from the client, write output,
// handle JSON, files, HTML templates, and manage redirects or errors in a structured way.
type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Params  map[string]string
}

// HTMLString writes the given HTML string directly to the response with the specified HTTP status code.
func (c *Context) HTMLString(status int, html_code string) {
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	c.Writer.WriteHeader(status)
	c.Writer.Write([]byte(html_code))
}

// Param returns the path parameter value for the given key.
func (c *Context) Param(key string) string {
	if c.Params == nil {
		return ""
	}
	return c.Params[key]
}

func (c *Context) ParamInt(key string) (int, error) {
	val := c.Param(key)
	return strconv.Atoi(val)
}

func (c *Context) ParamIntDefault(key string, defaultVal int) int {
	val := c.Param(key)
	if val == "" {
		return defaultVal
	}
	num, err := strconv.Atoi(val)
	if err != nil {
		LogError("ParamIntDefault", err.Error())
		return defaultVal
	}
	return num
}

// UserAgent returns the User-Agent header from the request.
func (c *Context) UserAgent() string {
	return c.Request.UserAgent()
}

// Query returns the query parameter value for the given key.
func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) QueryInt(key string) (int, error) {
	val := c.Query(key)
	return strconv.Atoi(val)
}

func (c *Context) QueryBool(key string, defaultVal bool) bool {
	val := c.Query(key)
	if val == "" {
		return defaultVal
	}

	lower := strings.ToLower(val)
	if lower == "true" || lower == "1" || lower == "yes" {
		return true
	} else if lower == "false" || lower == "0" || lower == "no" {
		return false
	}

	return defaultVal
}

func (c *Context) QueryIntDefault(key string, defaulVal int) int {
	val := c.Query(key)
	if val == "" {
		return defaulVal
	}
	num, err := strconv.Atoi(val)
	if err != nil {
		return defaulVal
	}
	return num
}

func (c *Context) QueryFloat(key string, defaultVal float64) float64 {
	val := c.Query(key)
	if val == "" {
		return defaultVal
	}
	num, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return defaultVal
	}
	return num
}

// It automatically hashes the data received from the form with Argon2.
func (c *Context) FormArgon2(key string) string {
	password := c.FormData(key) // formdan şifre al

	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		LogError("FormArrgon2()", err.Error())
		return ""
	}

	time := uint32(3)
	memory := uint32(64 * 1024)
	threads := uint8(4)
	keyLen := uint32(32)

	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	encodedHash := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		memory, time, threads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash))

	return encodedHash
}

// Automatically hashes the data received from the form with Bcrypt
func (c *Context) FormBcrypt(key string) string {
	password := c.FormData(key) // Formdan gelen veriyi al
	cost := 12
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		LogError("FormBcrypt()", err.Error())
		return ""
	}
	return string(hash)
}

// Redirect sends an HTTP redirect to the specified URL with the given status code.
// Commonly used status codes are 302 (Found), 301 (Moved Permanently), and 307/308.
func (c *Context) Redirect(status int, url string) {
	http.Redirect(c.Writer, c.Request, url, status)
}

// FormFile retrieves the uploaded file and its header from a multipart form
// for the given form key. It returns the file, its header, and an error if
// the file cannot be found or opened
func (c *Context) FormFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return c.Request.FormFile(key)
}

// FormData returns the value of a form field for the given key.
// It reads the data from POST or PUT form submissions.
func (c *Context) FormData(key string) string {
	return c.Request.FormValue(key)
}

// FormInt returns the value of a form field as an integer for the given key.
// If the conversion fails, it logs the error and returns 0.
func (c *Context) FormInt(key string) int {
	val := c.FormData(key)
	num, err := strconv.Atoi(val)
	if err != nil {
		LogError("FormInt()", err.Error())
		return 0
	}
	return num
}

// Hashes the data received from the form with MD5 and returns it.
func (c *Context) FormMD5(key string) string {
	hash := md5.Sum([]byte(c.FormData(key)))
	return hex.EncodeToString(hash[:])
}

// Hashes the data received from the form with Sha256 and returns it.
func (c *Context) FormSHA256(key string) string {
	hash := sha256.Sum256([]byte(c.FormData(key)))
	return hex.EncodeToString(hash[:])
}

// Hashes the data received from the form with Sha512 and returns it.
func (c *Context) FormSHA512(key string) string {
	hash := sha512.Sum512([]byte(c.FormData(key)))
	return hex.EncodeToString(hash[:])
}

// Delete returns true if the current request method is DELETE.
func (c *Context) Delete() bool {
	return c.Request.Method == http.MethodDelete
}

// Put returns true if the current request method is PUT.
func (c *Context) Put() bool {
	return c.Request.Method == http.MethodPut
}

// Post returns true if the current request method is POST.
func (c *Context) Post() bool {
	return c.Request.Method == http.MethodPost
}

// Get returns true if the current request method is GET.
func (c *Context) Get() bool {
	return c.Request.Method == http.MethodGet
}

// File serves a file from the given file path to the client.
func (c *Context) File(filepath string) {
	http.ServeFile(c.Writer, c.Request, filepath)
}

// JSONFile serves a JSON file from the given file path with the specified HTTP status code.
// Returns 404 if the file is not found.
func (c *Context) JSONFile(status int, filepath string) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		http.Error(c.Writer, "file not found", http.StatusNotFound)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Writer.WriteHeader(status)
	c.Writer.Write(data)
}

// JSONPretty encodes the given data as pretty-printed JSON and writes it to the response
// with the specified HTTP status code. Optional indent size can be provided.
func (c *Context) JSONPretty(status int, data interface{}, indent ...int) {
	ind := 2
	if len(indent) > 0 {
		ind = indent[0]
	}

	bytes, err := json.MarshalIndent(data, "", strings.Repeat(" ", ind))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	c.Writer.Write(bytes)
}

// JSON encodes the given data as JSON and writes it to the response with the specified HTTP status code.
func (c *Context) JSON(status int, data interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	json.NewEncoder(c.Writer).Encode(data)
}

// XMLPretty serializes the provided Go data structure into a pretty-printed XML
// and writes it as the response with indentation for readability.
func (c *Context) XMLPretty(status int, data interface{}) {
	c.Writer.Header().Set("Content-Type", "application/xml")
	c.Writer.WriteHeader(status)

	bytes, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Writer.Write([]byte(xml.Header))
	c.Writer.Write(bytes)
}

// XMLFile reads an XML file from the given filepath and writes it as the response.
// If the file cannot be found, it returns a 404 Not Found error.
// This is typically used for serving static XML files like sitemap.xml
func (c *Context) XMLFile(status int, filepath string) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		http.Error(c.Writer, "file not found", http.StatusNotFound)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/xml; charset=utf-8")
	c.Writer.WriteHeader(status)
	c.Writer.Write(data)
}

// XML serializes the provided Go data structure into XML and writes it as the response.
// The response Content-Type is set to application/xml.
// This is typically used for sending dynamic XML generated from structs.
func (c *Context) XML(status int, data interface{}) {
	c.Writer.Header().Set("Content-Type", "application/xml")
	c.Writer.WriteHeader(status)
	xml.NewEncoder(c.Writer).Encode(data)
}

// Stringf writes a formatted string to the response with the specified HTTP status code.
func (c *Context) Stringf(status int, text string, a ...any) {
	c.Writer.Header().Set("Content-Type", "text/plain")
	c.Writer.WriteHeader(status)
	c.Writer.Write([]byte(fmt.Sprintf(text, a...)))
}

// String writes a plain string to the response with the specified HTTP status code.
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

func (c *Context) Default405() {
	c.Writer.WriteHeader(http.StatusMethodNotAllowed)
	c.Writer.Write([]byte(`
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Method Not Allowed - 405 Error Flint</title>
    <style>
        :root {
            --primary-color: #218c87;
            --primary-dark: #1a6f6b;
            --primary-light: #e8f4f3;
            --text-color: #333;
            --white: #ffffff;
        }
        
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
        }
        
        body {
            background-color: #f9f9f9;
            color: var(--text-color);
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            height: 100vh;
            text-align: center;
            padding: 20px;
        }
        
        .container {
            max-width: 600px;
            width: 100%;
        }
        
        .error-image {
            width: 200px;
            height: 200px;
            margin-bottom: 30px;
            object-fit: contain;
        }
        
        h1 {
            font-size: 5rem;
            color: var(--primary-color);
            margin-bottom: 20px;
        }
        
        h2 {
            font-size: 1.8rem;
            margin-bottom: 15px;
            color: var(--text-color);
        }
        
        p {
            font-size: 1.1rem;
            margin-bottom: 30px;
            line-height: 1.6;
            color: #666;
        }
        
        .btn {
            display: inline-block;
            background-color: var(--primary-color);
            color: var(--white);
            padding: 12px 30px;
            border-radius: 50px;
            text-decoration: none;
            font-weight: 600;
            transition: all 0.3s ease;
            border: 2px solid var(--primary-color);
            margin: 0 10px;
        }
        
        .btn:hover {
            background-color: var(--primary-dark);
            border-color: var(--primary-dark);
            transform: translateY(-2px);
            box-shadow: 0 5px 15px rgba(33, 140, 135, 0.3);
        }
        
        .btn-outline {
            background-color: transparent;
            color: var(--primary-color);
        }
        
        .btn-outline:hover {
            background-color: var(--primary-light);
            color: var(--primary-dark);
        }
        
        .animation {
            animation: float 3s ease-in-out infinite;
        }
        
        @keyframes float {
            0%, 100% {
                transform: translateY(0);
            }
            50% {
                transform: translateY(-15px);
            }
        }
        
        @media (max-width: 768px) {
            h1 {
                font-size: 3.5rem;
            }
            
            h2 {
                font-size: 1.4rem;
            }
            
            .btn {
                display: block;
                margin: 10px auto;
                width: 80%;
                max-width: 250px;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <!-- Add your own image here -->
        <h1>Flint Framework</h1>
        <h1>405</h1>
        <h2>Method not allowed</h2>
        <p>This endpoint does not support this HTTP method. Please try another action.</p>
        
        <div class="buttons">
            <a href="/" class="btn">Return to Homepage</a>
        </div>
    </div>

    <script>
        // Simple animation effects
        document.addEventListener('DOMContentLoaded', function() {
            const image = document.querySelector('.error-image');
            
            // Pause animation on hover
            image.addEventListener('mouseenter', function() {
                this.style.animationPlayState = 'paused';
            });
            
            // Resume animation when mouse leaves
            image.addEventListener('mouseleave', function() {
                this.style.animationPlayState = 'running';
            });
            
            // Button click effect
            const buttons = document.querySelectorAll('.btn');
            buttons.forEach(button => {
                button.addEventListener('click', function() {
                    this.style.transform = 'scale(0.95)';
                    setTimeout(() => {
                        this.style.transform = 'scale(1)';
                    }, 200);
                });
            });
        });
    </script>
</body>
</html>
    `))
}

func (c *Context) Default404() {
	c.Writer.WriteHeader(http.StatusNotFound)
	c.Writer.Write([]byte(`
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Page Not Found - 404 Error Flint</title>
    <style>
        :root {
            --primary-color: #218c87;
            --primary-dark: #1a6f6b;
            --primary-light: #e8f4f3;
            --text-color: #333;
            --white: #ffffff;
        }
        
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
        }
        
        body {
            background-color: #f9f9f9;
            color: var(--text-color);
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            height: 100vh;
            text-align: center;
            padding: 20px;
        }
        
        .container {
            max-width: 600px;
            width: 100%;
        }
        
        .error-image {
            width: 200px;
            height: 200px;
            margin-bottom: 30px;
            object-fit: contain;
        }
        
        h1 {
            font-size: 5rem;
            color: var(--primary-color);
            margin-bottom: 20px;
        }
        
        h2 {
            font-size: 1.8rem;
            margin-bottom: 15px;
            color: var(--text-color);
        }
        
        p {
            font-size: 1.1rem;
            margin-bottom: 30px;
            line-height: 1.6;
            color: #666;
        }
        
        .btn {
            display: inline-block;
            background-color: var(--primary-color);
            color: var(--white);
            padding: 12px 30px;
            border-radius: 50px;
            text-decoration: none;
            font-weight: 600;
            transition: all 0.3s ease;
            border: 2px solid var(--primary-color);
            margin: 0 10px;
        }
        
        .btn:hover {
            background-color: var(--primary-dark);
            border-color: var(--primary-dark);
            transform: translateY(-2px);
            box-shadow: 0 5px 15px rgba(33, 140, 135, 0.3);
        }
        
        .btn-outline {
            background-color: transparent;
            color: var(--primary-color);
        }
        
        .btn-outline:hover {
            background-color: var(--primary-light);
            color: var(--primary-dark);
        }
        
        .animation {
            animation: float 3s ease-in-out infinite;
        }
        
        @keyframes float {
            0%, 100% {
                transform: translateY(0);
            }
            50% {
                transform: translateY(-15px);
            }
        }
        
        @media (max-width: 768px) {
            h1 {
                font-size: 3.5rem;
            }
            
            h2 {
                font-size: 1.4rem;
            }
            
            .btn {
                display: block;
                margin: 10px auto;
                width: 80%;
                max-width: 250px;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <!-- Add your own image here -->
        <h1>Flint Framework Official</h1>
        <h1>404</h1>
        <h2>Page Not Found</h2>
        <p>The page you are looking for might have been removed, had its name changed, or is temporarily unavailable. Please try to find what you need from our homepage.</p>
        
        <div class="buttons">
            <a href="/" class="btn">Return to Homepage</a>
        </div>
    </div>

    <script>
        // Simple animation effects
        document.addEventListener('DOMContentLoaded', function() {
            const image = document.querySelector('.error-image');
            
            // Pause animation on hover
            image.addEventListener('mouseenter', function() {
                this.style.animationPlayState = 'paused';
            });
            
            // Resume animation when mouse leaves
            image.addEventListener('mouseleave', function() {
                this.style.animationPlayState = 'running';
            });
            
            // Button click effect
            const buttons = document.querySelectorAll('.btn');
            buttons.forEach(button => {
                button.addEventListener('click', function() {
                    this.style.transform = 'scale(0.95)';
                    setTimeout(() => {
                        this.style.transform = 'scale(1)';
                    }, 200);
                });
            });
        });
    </script>
</body>
</html>
	`))
}

func (c *Context) Template405(templatePath string, data ...any) {
	var d any
	if len(data) > 0 {
		d = data[0]
	} else {
		d = map[string]any{}
	}
	c.HTML(http.StatusMethodNotAllowed, templatePath, d)
}

func (c *Context) Template404(templatePath string, data ...any) {
	var d any
	if len(data) > 0 {
		d = data[0]
	} else {
		d = map[string]any{}
	}
	c.HTML(http.StatusNotFound, templatePath, d)
}
