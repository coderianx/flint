# Using Query Parameters

Flint Framework provides several methods to easily retrieve query parameters from the URL.  
These methods allow you to read parameters as **string**, **int**, or **float** types, and return default values in case of errors.

---

## Available Methods

### `Query(key string) string`
Returns the parameter as a string.  
If the parameter does not exist, it returns an empty string.

### `QueryInt(key string) (int, error)`
Returns the parameter as an integer.  
The return values are the integer value and an `error`.  
If the parameter is not a number, it returns an error.

### `QueryIntDefault(key string, defaultVal int) int`
Returns the parameter as an integer.  
If the parameter does not exist or is not a valid number, the **default** value is returned.

### `QueryFloat(key string, defaultVal float64) float64`
Returns the parameter as a float64.  
If the parameter does not exist or is not a valid number, the **default** value is returned.

---

## Example Usage

```go
package main

import (
	"fmt"
	"github.com/coderianx/flint"
)

func main() {
	app := flint.NewServer()

	// /search endpoint
	// Example URL: http://localhost:8080/search?q=golang&page=2&rating=4.5
	app.Handle("/search", func(ctx *flint.Context) {
		// String query param
		query := ctx.Query("q") // returns string

		// Int query param (error-returning version)
		page, err := ctx.QueryInt("page")
		if err != nil {
			page = 1 // default value
		}

		// Int query param (default value version)
		limit := ctx.QueryIntDefault("limit", 10)

		// Float query param (default value version)
		rating := ctx.QueryFloat("rating", 0.0)

		ctx.JSON(200, map[string]interface{}{
			"query":  query,
			"page":   page,
			"limit":  limit,
			"rating": rating,
		})
	})

	app.Run(":8080")
}
```

---

## Testing

Run the server:
```bash
go run main.go
```

Make a request via browser or cURL:
```bash
curl "http://localhost:8080/search?q=flint&page=2&limit=20&rating=4.8"
```

### Expected JSON Output
```json
{
  "query": "flint",
  "page": 2,
  "limit": 20,
  "rating": 4.8
}
```

---

## Notes
- Always check for errors when using `QueryInt`.  
- `QueryIntDefault` and `QueryFloat` methods automatically return the default value when the parameter is missing or invalid.
