package main

import "github.com/coderianx/flint"

// TR: Basit bir örnek REST API
// Bu örnekte /users adresine gittiğinizde JSON formatında kullanıcı listesi döndürülür.
//
// EN: Simple example REST API
// In this example, visiting /users will return a JSON-formatted list of users.
func main() {
	app := flint.NewServer()

	// TR: JSON olarak döndürülecek kullanıcı listesi
	// EN: User list to be returned as JSON
	users := []map[string]interface{}{
		{"id": 1, "name": "Flint"},
		{"id": 2, "name": "Framework"},
	}

	// TR: /users isteği geldiğinde kullanıcı listesini JSON olarak döndür
	// EN: When /users is requested, return the user list as JSON
	app.Handle("/users", func(ctx *flint.Context) {
		ctx.JSON(200, users)
	})

	// TR: Sunucuyu 8080 portunda başlat
	// EN: Start the server on port 8080
	app.Run(":8080")
}
