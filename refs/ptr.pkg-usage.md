# ptr.pkg-usage

## Card

Why imported structs (like database connections, HTTP clients, configurations) are typed/passed as pointers.

- **Prevent Copying Stateful Objects:** Stateless code is fine, but stateful objects (with mutexes, files, connection pools) must never be copied.
- **Shared State:** Ensures all parts of your app interact with the same underlying state.

### Subcard: Stateful Struct Safety

```go
type Server struct {
	mu     sync.Mutex // Mutex MUST NOT be copied!
	Active bool
}
// MUST use pointer receivers to avoid copying sync.Mutex
func (s *Server) Start() {}
```

## Example

Creating and passing a configuration pointer across packages:

```go
package main

import (
	"database/sql"
	"net/http"
)

type App struct {
	DB     *sql.DB      // Pointer to shared database pool
	Client *http.Client // Pointer to shared HTTP client
}

func main() {
	// sql.Open returns a pointer to sql.DB
	db, _ := sql.Open("postgres", "dsn")
	
	app := &App{
		DB:     db,
		Client: http.DefaultClient,
	}
	
	_ = app
}
```

## Deep

Many structures in imported packages are "stateful" or contain unexported internal state:
1. **Mutexes and Sync Primitives:** Structs like `sync.Mutex` or `sync.WaitGroup` contain internal state tracking locks. Copying a mutex copy-locks it, leading to deadlocks or runtime panics.
2. **Resource Handles:** Structures representing databases (`sql.DB`), network connections, or files contain active operating system descriptors. If you copy them, the copies share descriptors but can close them out of order, corrupting resources.
3. **Large Objects:** Libraries pass pointers to avoid expensive $O(N)$ memory allocations.

## Gotchas

- **Returning Pointers:** Constructors in packages (like `NewClient()`) almost always return a pointer (e.g., `*Client`) rather than a value, indicating that the struct is stateful and should be managed as a reference.

## Related

- ptr.basics
- struct.methods
- ptr.flow-chart
