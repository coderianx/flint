package main

import (
	"strconv"

	"github.com/coderianx/flint"
)

// TR: User yapısı - her kullanıcı bir ID ve isim içerir
// EN: User struct - each user has an ID and a name
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	app := flint.NewServer()

	// TR: Örnek kullanıcı listesi
	// EN: Example user list
	users := []User{
		{ID: 1, Name: "Flint"},
		{ID: 2, Name: "Framework"},
		{ID: 3, Name: "GoLang"},
	}

	// TR: /users endpoint'i - tüm kullanıcılar veya belirli bir kullanıcıyı döndürür
	// EN: /users endpoint - returns all users or a specific user
	app.Handle("/users", func(ctx *flint.Context) {
		if ctx.Get() {
			// TR: Query parametre kontrolü - ?id=1 şeklinde
			// EN: Check query parameter - like ?id=1
			idStr := ctx.Query("id")

			if idStr == "" {
				// TR: id verilmemişse tüm kullanıcıları döndür
				// EN: If no id is provided, return all users
				ctx.JSON(200, users)
				return
			}

			// TR: id parametresini int'e çevir
			// EN: Convert id parameter to int
			id, err := strconv.Atoi(idStr)
			if err != nil {
				// TR: id geçersizse hata döndür
				// EN: Return error if id is invalid
				ctx.JSON(400, map[string]string{"error": "Invalid ID"})
				return
			}

			// TR: Kullanıcıyı listede ara
			// EN: Search for the user in the list
			for _, u := range users {
				if u.ID == id {
					// TR: Kullanıcı bulunduğunda döndür
					// EN: Return the user if found
					ctx.JSON(200, u)
					return
				}
			}

			// TR: Kullanıcı bulunamadığında hata döndür
			// EN: Return error if user not found
			ctx.JSON(404, map[string]string{"error": "User not found"})
		} else {
			// TR: GET dışında istek yapılırsa 405 döndür
			// EN: Return 405 if method is not GET
			ctx.String(405, "405 Method Not Allowed")
		}
	})

	// TR: Sunucuyu 8080 portunda başlat
	// EN: Start the server on port 8080
	app.Run(":8080")
}
