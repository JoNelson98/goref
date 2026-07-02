# http.query-params

## Card

Parse URL query parameters (e.g. `/todos?completed=true`) using `r.URL.Query()`.

```go
queryParams := r.URL.Query()
completed := queryParams.Get("completed") // returns "" if not found
```

### Subcard: Simple Check

```go
if r.URL.Query().Get("important") == "true" {
    // filter important items
}
```

## Example

```go
package main

import (
	"fmt"
	"net/http"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	// Query is of type url.Values, which is map[string][]string
	query := r.URL.Query()
	
	searchTerm := query.Get("q") // Gets the first value for key "q"
	limitStr := query.Get("limit")

	fmt.Fprintf(w, "Searching for: %q, Limit: %q\n", searchTerm, limitStr)
}

func main() {
	http.HandleFunc("/search", searchHandler)
	// Try: http://localhost:8080/search?q=gopher&limit=10
}
```

## Deep

`r.URL.Query()` parses the raw URL query string and returns a `url.Values` map.
- `url.Values` is a map of string keys to slices of strings: `map[string][]string`. This is because a query key can be defined multiple times (e.g. `?tags=go&tags=web`).
- The `.Get(key)` method returns the *first* value associated with the given key. If the key is not present, it returns an empty string `""`.
- To get all values of a multi-valued query parameter, access the map slice directly: `values := r.URL.Query()["tags"]`.

## Gotchas

- **Get always returns string:** `.Get(key)` returns a `string`. If you are expecting an integer (like a page limit or boolean completed flag), you must parse it yourself using standard functions like `strconv.Atoi` or `strconv.ParseBool`.

## Related

- http.crud-handler
