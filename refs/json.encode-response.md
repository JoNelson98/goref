# json.encode-response

## Card

Serialize a Go struct or map into a JSON string and write it to an HTTP response stream using `json.NewEncoder`.

```go
w.Header().Set("Content-Type", "application/json")
err := json.NewEncoder(w).Encode(payload)
```

### Subcard: Simple Marshal to Bytes

For in-memory JSON serialization (not streaming):
```go
bytes, err := json.Marshal(payload)
```

## Example

```go
package main

import (
	"encoding/json"
	"net/http"
)

type UserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Set response header to JSON
	w.Header().Set("Content-Type", "application/json")

	resp := UserResponse{
		ID:    42,
		Email: "gopher@golang.org",
	}

	// 2. Stream JSON directly to ResponseWriter
	w.WriteHeader(http.StatusOK) // 200 OK
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
```

## Deep

Using `json.NewEncoder(w).Encode(data)` is highly optimized for network requests. It streams JSON bytes directly to the socket connection buffer (`io.Writer`), which avoids allocating a temporary slice of bytes in-memory—unlike `json.Marshal(data)` which allocates a full byte slice (`[]byte`) representing the entire JSON payload in RAM before you write it.

## Gotchas

- **Header Order:** Always set your headers (e.g. `w.Header().Set("Content-Type", ...)`) **before** calling `w.WriteHeader` or writing any body bytes. Once you call `w.WriteHeader` or write to the body, headers are locked and sent to the client; trying to modify them afterwards will be ignored silently.
- **HTML Escaping:** By default, Go's JSON encoder escapes characters like `<` and `>` into their Unicode representations to prevent cross-site scripting (XSS). To disable this, use a custom encoder:
  ```go
  enc := json.NewEncoder(w)
  enc.SetEscapeHTML(false)
  ```

## Related

- json.decode-request
- http.crud-handler
- struct.basics
