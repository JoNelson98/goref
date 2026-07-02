# json.decode-request

## Card

Decode JSON from an HTTP request body using `json.NewDecoder`.

```go
var payload MyStruct
err := json.NewDecoder(r.Body).Decode(&payload)
```

### Subcard: Strict Decoding

Reject unknown JSON fields using `DisallowUnknownFields`:
```go
dec := json.NewDecoder(r.Body)
dec.DisallowUnknownFields()
err := dec.Decode(&payload)
```

## Example

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Received: %+v\n", u)
}

func main() {
	body := strings.NewReader(`{"name": "Alice", "age": 30}`)
	req, _ := http.NewRequest("POST", "/user", body)
	
	// Simulated response recorder
	// http.HandlerFunc(handleUser).ServeHTTP(rec, req)
	_ = req
}
```

## Deep

For decoding data from stream sources like connection sockets, files, or `http.Request.Body`, using `json.NewDecoder` is more memory efficient than reading the entire body into a byte slice with `io.ReadAll` and then running `json.Unmarshal`. `NewDecoder` decodes chunks directly from the reader.

## Gotchas

- Don't forget to close `r.Body` (though the HTTP server usually handles this for the main request body, custom clients or roundtrippers require manual closing).
- `Decode` might return `io.EOF` if the body is empty.

## Related

- http.make-api-call
