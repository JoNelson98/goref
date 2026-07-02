# http.server

## Card

Start a standard library HTTP server in Go on a specified port.

```go
http.HandleFunc("/route", handler)
err := http.ListenAndServe(":8080", nil)
```

### Subcard: Start with Custom Server

```go
srv := &http.Server{
	Addr:         ":8080",
	ReadTimeout:  10 * time.Second,
	WriteTimeout: 10 * time.Second,
}
err := srv.ListenAndServe()
```

## Example

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go Web Server!")
}

func main() {
	// 1. Register endpoint route
	http.HandleFunc("/", helloHandler)

	fmt.Println("Server starting on port 8080...")

	// 2. Define custom server with timeout settings for production
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// 3. Start listener (blocks until server stops)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server stopped with error:", err)
	}
}
```

## Deep

`http.ListenAndServe` starts a server and blocks the current thread until the program is shut down or a fatal error occurs:
- The route registration `http.HandleFunc` maps paths to functions in the `http.DefaultServeMux` (Go's built-in multiplexer).
- Handlers are executed concurrently. Every incoming connection spawns a new goroutine to invoke the handler function, meaning state shared across routes must be lock-protected.

## Gotchas

- **Production Timeout Safety:** Never call `http.ListenAndServe` directly in a production environment. By default, it uses no read, write, or idle timeouts. This makes your server highly vulnerable to Slowloris attacks or socket depletion (leaking unclosed connection sockets). Always build a custom `&http.Server{}` with explicit timeouts.
- **Multiplexer Collision:** Standard library router does exact path matches. Be careful when matching `/` as it behaves as a catch-all route for any undefined paths.

## Related

- http.crud-handler
- os.getenv
