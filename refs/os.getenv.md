# os.getenv

## Card

Read environmental configuration variables using the standard library `os` package.

```go
port := os.Getenv("PORT") // Returns "" if not set
```

### Subcard: Lookup Environment

```go
val, exists := os.LookupEnv("DB_URL")
if !exists {
    // Variable is not set
}
```

## Example

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. Get an existing or fallback environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Fallback default
	}
	fmt.Println("Server will bind to port:", port)

	// 2. Strict lookup (distinguishes empty vs unset)
	dbURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		fmt.Println("DATABASE_URL is not set!")
	} else {
		fmt.Printf("Database target: %s\n", dbURL)
	}
}
```

## Deep

Twelve-Factor App guidelines advocate storing application configuration in environment variables (ports, keys, database connection strings). 
- `os.Getenv(key)` fetches values directly from the host system environment.
- It returns an empty string `""` if the variable is not set.
- To distinguish between "variable is empty string" and "variable is completely unset", use `os.LookupEnv(key)`.

## Gotchas

- **No Type Casting:** Environmental variables are always returned as strings. If you need integer types (e.g. `PORT`), parse them yourself:
  ```go
  port, err := strconv.Atoi(os.Getenv("PORT"))
  ```
- **Loading .env files:** The Go standard library does **not** read `.env` files automatically. In production, configure environment variables in your deployment wrapper (Docker, Kubernetes). In development, load them manually or use libraries like `github.com/joho/godotenv` to populate them.

## Related

- http.server
