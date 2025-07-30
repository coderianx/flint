package flint

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"html/template"
	"mime/multipart"
	"net/http"
	"strconv"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (c *Context) UserAgent() string {
	return c.Request.UserAgent()
}

func (c *Context) Send(html string) {
	c.ResponseWriter.Header().Set("Content-Type", "text/html")
	c.ResponseWriter.WriteHeader(200)
	c.ResponseWriter.Write([]byte(html))
}

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

func (c *Context) Redirect(status int, url string) {
	http.Redirect(c.Writer, c.Request, url, status)
}

func (c *Context) FormFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return c.Request.FormFile(key)
}

func (c *Context) FormData(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) FormInt(key string) int {
	val := c.FormData(key)
	num, err := strconv.Atoi(val)
	if err != nil {
		LogError("FormInt()", err.Error())
		return 0
	}
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

func (c *Context) Delete() bool {
	return c.Request.Method == http.MethodDelete
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
        <img src="assets/flint.jpg" alt="404 Error" class="error-image animation">
        
        <h1>404</h1>
        <h2>Page Not Found</h2>
        <p>The page you are looking for might have been removed, had its name changed, or is temporarily unavailable. Please try to find what you need from our homepage.</p>
        
        <div class="buttons">
            <a href="/" class="btn">Return to Homepage</a>
            <a href="https://t.me/Grayvort3x" class="btn btn-outline">Contact Support</a>
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

func (c *Context) Template404(templatePath string, data ...any) {
	var d any
	if len(data) > 0 {
		d = data[0]
	} else {
		d = map[string]any{}
	}
	c.HTML(http.StatusNotFound, templatePath, d)
}
