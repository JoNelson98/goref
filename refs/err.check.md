# err.check

## Card

Go treats errors as normal values returned from functions. Always check the error value immediately.

```go
val, err := doSomething()
if err != nil {
    // Handle the error (log, return, or exit)
}
```

### Subcard: Custom Errors

Create a simple custom error using `errors.New` or `fmt.Errorf`:
```go
err := errors.New("something went wrong")
errWithCtx := fmt.Errorf("failed to process user %d: %w", userID, err)
```

## Example

```go
package main

import (
	"errors"
	"fmt"
)

// Define package-level sentinel errors for callers to check
var ErrItemNotFound = errors.New("item not found in store")

func findItem(id int) (string, error) {
	if id < 0 {
		return "", errors.New("invalid item ID") // Simple custom error
	}
	if id == 42 {
		return "Golden Ticket", nil
	}
	return "", ErrItemNotFound
}

func main() {
	item, err := findItem(-1)
	if err != nil {
		fmt.Println("Error occurred:", err) // Error occurred: invalid item ID
	}

	_, err = findItem(100)
	if errors.Is(err, ErrItemNotFound) {
		fmt.Println("Item was missing!") // Item was missing!
	}
}
```

## Deep

In Go, error handling is explicit rather than implicit (no try-catch blocks):
- **Sentinel Errors:** Package-level constants (like `io.EOF` or custom `ErrItemNotFound`) used for standard condition checking. Use `errors.Is(err, ErrItem)` to check if a wrapped error matches a sentinel error.
- **Custom Error Types:** You can define structs that implement the `error` interface (which requires a single method `Error() string`) to pass rich metadata context:
  ```go
  type QueryError struct { Query string; Err error }
  func (e *QueryError) Error() string { return fmt.Sprintf("query %q failed: %v", eQuery, e.Err) }
  ```

## Gotchas

- **Do not ignore errors:** Never use the blank identifier `_` to discard errors returned by functions unless it is a safe operation where errors are mathematically impossible. Ignoring errors leads to runtime panics and corrupt states.
- **Wrapped Errors Checking:** Don't check wrapped errors using simple string equality (e.g. `err.Error() == "not found"`). Always use `errors.Is` (for comparing error values) or `errors.As` (for checking custom error struct types) so wrapper chains are traversed correctly.

## Related

- ptr.nil-checks
