# User Agent 

Here is an example of how you can get User Agent information with **flint framework**.

## User Agent Information Retrieval

User Agent is a header sent by the client (browser, application, etc.) to the server. You can access User Agent information in Flint using the `ctx.UserAgent()` method.

```go
package main

import (
    "fmt"
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        if ctx.Get() {
            useragent := ctx.UserAgent() // Get User Agent information
            fmt.Printf("User Agent: %s\n", useragent)
        }
    })
}
```
**Explanation:**
- `ctx.UserAgent()`: Retrieves the User Agent information sent by the client.
- `fmt.Printf("User Agent: %s\n", useragent)`: Prints the User Agent information to the console.