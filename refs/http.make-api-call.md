# http.make-api-call

## Card

Make an HTTP GET or POST call using `http.Client` or package helpers.

```go
// Simple GET
resp, err := http.Get("https://api.example.com")

// Custom Request (needed for custom headers, e.g. Auth)
req, err := http.NewRequest(http.MethodGet, url, nil)
req.Header.Set("Authorization", "Bearer token")
resp, err := http.DefaultClient.Do(req)
```

### Subcard: Simple GET Request

```go
resp, err := http.Get("https://api.github.com")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)
```

### Subcard: POST JSON Request

```go
payload := strings.NewReader(`{"name":"test"}`)
resp, err := http.Post("https://api.example.com", "application/json", payload)
```

## Example

```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}
	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("Body length:", len(body))
}
```

## Deep

Always use `http.NewRequest` if you need to pass query parameters, set custom headers, or manage request contexts. Using context (`http.NewRequestWithContext`) is standard in production environments to support timeouts and cancellations.

## Gotchas

- **Resource Leak:** You MUST close `resp.Body` if `err == nil`. If you don't close it, the underlying TCP connection remains open/allocated and will eventually exhaust system sockets.
- `http.Get` or `http.Post` use `http.DefaultClient` which has NO timeout. In production, always define a custom `http.Client` with a timeout:
  ```go
  var client = &http.Client{
      Timeout: 15 * time.Second,
  }
  ```

## Related

- json.decode-request
