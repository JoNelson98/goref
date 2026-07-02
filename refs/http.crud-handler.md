# http.crud-handler

## Card

Standard library `net/http` in-memory CRUD REST API for a Todo application.

```go
func handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":    // Read All
	case "POST":   // Create
	case "PUT":    // Update (on /todos/<id>)
	case "DELETE": // Delete (on /todos/<id>)
	default:       http.Error(w, "Method not allowed", 405)
	}
}
```

### Subcard: Define Data Struct

```go
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
```

## Example

Complete working in-memory CRUD server for Todos using only standard library `net/http`:

```go
package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var (
	todos  = []Todo{}
	nextID = 1
	mu     sync.Mutex
)

func handleTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	switch r.Method {
	case "GET":
		mu.Lock()
		json.NewEncoder(w).Encode(todos)
		mu.Unlock()

	case "POST":
		var t Todo
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		mu.Lock()
		t.ID = nextID
		nextID++
		todos = append(todos, t)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(t)
		mu.Unlock()

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Handler for a specific todo: /todos/1
func handleTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "PUT":
		var t Todo
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		mu.Lock()
		defer mu.Unlock()
		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Title = t.Title
				todos[i].Completed = t.Completed
				json.NewEncoder(w).Encode(todos[i])
				return
			}
		}
		http.Error(w, "Todo not found", http.StatusNotFound)

	case "DELETE":
		mu.Lock()
		defer mu.Unlock()
		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		http.Error(w, "Todo not found", http.StatusNotFound)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/todos", handleTodos)
	http.HandleFunc("/todos/", handleTodo) // Matches /todos/<id>
	http.ListenAndServe(":8080", nil)
}
```

## Deep

To build a REST API in Go without third-party frameworks:
- Use `http.HandleFunc("/todos", handleTodos)` to handle collection routes.
- Use `http.HandleFunc("/todos/", handleTodo)` with a trailing slash to handle item routes like `/todos/1`. The trailing slash acts as a wildcard, catching any subpaths.
- Use a `sync.Mutex` (`mu`) when reading or writing to global shared maps or slices. Since HTTP handlers in Go execute concurrently in separate goroutines, a mutex prevents **concurrent map writes** or data races that would crash your server.

## Gotchas

- **Trailing Slash Wildcard:** `http.HandleFunc("/todos", ...)` will strictly match `/todos` only. `http.HandleFunc("/todos/", ...)` will match `/todos/` and any path starting with `/todos/` (like `/todos/123`).
- **Concurrent Request Safety:** Slices are not thread-safe. Always protect read/write state operations with a Mutex (`sync.Mutex`) in standard HTTP handlers.

## Related

- json.decode-request
- slice.remove-item
- struct.basics
