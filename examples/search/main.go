package main

import (
	"strings"

	"github.com/coderianx/flint"
)

// Article struct — Makale verilerini temsil eden yapı
// This struct represents an article with ID, Title, and Content fields.
type Article struct {
	ID      int    `json:"id"`      // Makale ID'si — Article ID
	Title   string `json:"title"`   // Makale başlığı — Article title
	Content string `json:"content"` // Makale içeriği — Article content
}

func main() {
	// Yeni bir Flint sunucusu başlat
	// Start a new Flint server
	app := flint.NewServer()

	// Örnek makale verileri
	// Example article data
	articles := []Article{
		{ID: 1, Title: "Golang ile Web 1", Content: "Go dilinde web geliştirme örnekleri 1..."},
		{ID: 2, Title: "Flint Framework Tanıtımı 2", Content: "Flint, hafif bir Go web frameworküdür 2..."},
		{ID: 3, Title: "JavaScript Temelleri 3", Content: "JavaScript değişkenler, fonksiyonlar 3..."},
		{ID: 4, Title: "Go Routines ve Concurrency 4", Content: "Go rutinleri ile paralel işler yapma 4..."},
		{ID: 5, Title: "Web için HTML & CSS 5", Content: "Temel HTML ve CSS stilleri 5..."},
		{ID: 6, Title: "Golang ile Web 6", Content: "Go dilinde web geliştirme örnekleri 6..."},
		{ID: 7, Title: "Flint Framework Tanıtımı 7", Content: "Flint, hafif bir Go web frameworküdür 7..."},
		{ID: 8, Title: "JavaScript Temelleri 8", Content: "JavaScript değişkenler, fonksiyonlar 8..."},
		{ID: 9, Title: "Go Routines ve Concurrency 9", Content: "Go rutinleri ile paralel işler yapma 9..."},
		{ID: 10, Title: "Web için HTML & CSS 10", Content: "Temel HTML ve CSS stilleri 10..."},
		{ID: 11, Title: "Golang ile Web 11", Content: "Go dilinde web geliştirme örnekleri 11..."},
		{ID: 12, Title: "Flint Framework Tanıtımı 12", Content: "Flint, hafif bir Go web frameworküdür 12..."},
		{ID: 13, Title: "JavaScript Temelleri 13", Content: "JavaScript değişkenler, fonksiyonlar 13..."},
		{ID: 14, Title: "Go Routines ve Concurrency 14", Content: "Go rutinleri ile paralel işler yapma 14..."},
		{ID: 15, Title: "Web için HTML & CSS 15", Content: "Temel HTML ve CSS stilleri 15..."},
		{ID: 16, Title: "Golang ile Web 16", Content: "Go dilinde web geliştirme örnekleri 16..."},
		{ID: 17, Title: "Flint Framework Tanıtımı 17", Content: "Flint, hafif bir Go web frameworküdür 17..."},
		{ID: 18, Title: "JavaScript Temelleri 18", Content: "JavaScript değişkenler, fonksiyonlar 18..."},
		{ID: 19, Title: "Go Routines ve Concurrency 19", Content: "Go rutinleri ile paralel işler yapma 19..."},
		{ID: 20, Title: "Web için HTML & CSS 20", Content: "Temel HTML ve CSS stilleri 20..."},
		{ID: 21, Title: "Golang ile Web 21", Content: "Go dilinde web geliştirme örnekleri 21..."},
		{ID: 22, Title: "Flint Framework Tanıtımı 22", Content: "Flint, hafif bir Go web frameworküdür 22..."},
		{ID: 23, Title: "JavaScript Temelleri 23", Content: "JavaScript değişkenler, fonksiyonlar 23..."},
		{ID: 24, Title: "Go Routines ve Concurrency 24", Content: "Go rutinleri ile paralel işler yapma 24..."},
		{ID: 25, Title: "Web için HTML & CSS 25", Content: "Temel HTML ve CSS stilleri 25..."},
		{ID: 26, Title: "Golang ile Web 26", Content: "Go dilinde web geliştirme örnekleri 26..."},
		{ID: 27, Title: "Flint Framework Tanıtımı 27", Content: "Flint, hafif bir Go web frameworküdür 27..."},
		{ID: 28, Title: "JavaScript Temelleri 28", Content: "JavaScript değişkenler, fonksiyonlar 28..."},
		{ID: 29, Title: "Go Routines ve Concurrency 29", Content: "Go rutinleri ile paralel işler yapma 29..."},
		{ID: 30, Title: "Web için HTML & CSS 30", Content: "Temel HTML ve CSS stilleri 30..."},
		{ID: 31, Title: "Golang ile Web 31", Content: "Go dilinde web geliştirme örnekleri 31..."},
		{ID: 32, Title: "Flint Framework Tanıtımı 32", Content: "Flint, hafif bir Go web frameworküdür 32..."},
		{ID: 33, Title: "JavaScript Temelleri 33", Content: "JavaScript değişkenler, fonksiyonlar 33..."},
		{ID: 34, Title: "Go Routines ve Concurrency 34", Content: "Go rutinleri ile paralel işler yapma 34..."},
		{ID: 35, Title: "Web için HTML & CSS 35", Content: "Temel HTML ve CSS stilleri 35..."},
		{ID: 36, Title: "Golang ile Web 36", Content: "Go dilinde web geliştirme örnekleri 36..."},
		{ID: 37, Title: "Flint Framework Tanıtımı 37", Content: "Flint, hafif bir Go web frameworküdür 37..."},
		{ID: 38, Title: "JavaScript Temelleri 38", Content: "JavaScript değişkenler, fonksiyonlar 38..."},
		{ID: 39, Title: "Go Routines ve Concurrency 39", Content: "Go rutinleri ile paralel işler yapma 39..."},
		{ID: 40, Title: "Web için HTML & CSS 40", Content: "Temel HTML ve CSS stilleri 40..."},
		{ID: 41, Title: "Golang ile Web 41", Content: "Go dilinde web geliştirme örnekleri 41..."},
		{ID: 42, Title: "Flint Framework Tanıtımı 42", Content: "Flint, hafif bir Go web frameworküdür 42..."},
		{ID: 43, Title: "JavaScript Temelleri 43", Content: "JavaScript değişkenler, fonksiyonlar 43..."},
		{ID: 44, Title: "Go Routines ve Concurrency 44", Content: "Go rutinleri ile paralel işler yapma 44..."},
		{ID: 45, Title: "Web için HTML & CSS 45", Content: "Temel HTML ve CSS stilleri 45..."},
		{ID: 46, Title: "Golang ile Web 46", Content: "Go dilinde web geliştirme örnekleri 46..."},
		{ID: 47, Title: "Flint Framework Tanıtımı 47", Content: "Flint, hafif bir Go web frameworküdür 47..."},
		{ID: 48, Title: "JavaScript Temelleri 48", Content: "JavaScript değişkenler, fonksiyonlar 48..."},
		{ID: 49, Title: "Go Routines ve Concurrency 49", Content: "Go rutinleri ile paralel işler yapma 49..."},
		{ID: 50, Title: "Web için HTML & CSS 50", Content: "Temel HTML ve CSS stilleri 50..."},
		// ...
	}

	// Ana sayfada statik HTML dosyasını servis et
	// Serve the static HTML file at the root path
	app.Handle("/", func(ctx *flint.Context) {
		ctx.File("./examples/search/index.html")
	})

	// Arama API endpoint'i
	// Search API endpoint
	app.Handle("/api/search", func(ctx *flint.Context) {
		// Sadece GET isteğine izin ver
		// Allow only GET requests
		if !ctx.Get() {
			ctx.String(405, "Method Not Allowed") // 405 — Method Not Allowed
			return
		}

		// Query parametre "q" değerini al ve küçük harfe çevir
		// Get the "q" query parameter and convert it to lowercase
		q := strings.ToLower(strings.TrimSpace(ctx.Query("q")))

		// Arama sonuçlarını tutmak için slice
		// Slice to store search results
		var results []Article

		// Eğer arama kelimesi boş değilse filtreleme yap
		// If search query is not empty, filter the articles
		if q != "" {
			for _, a := range articles {
				// Başlıkta veya içerikte arama kelimesi geçiyorsa ekle
				// If the search query appears in the title or content, add it to results
				if strings.Contains(strings.ToLower(a.Title), q) ||
					strings.Contains(strings.ToLower(a.Content), q) {
					results = append(results, a)
				}
			}
		}

		// JSON formatında sonuçları döndür
		// Return results in JSON format
		ctx.JSON(200, map[string]interface{}{
			"count":   len(results), // Toplam bulunan sonuç sayısı — Total number of results
			"results": results,      // Bulunan makaleler — Found articles
		})
	})

	// Sunucuyu 8080 portunda çalıştır
	// Run the server on port 8080
	app.Run(":8080")
}
