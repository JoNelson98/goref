# cmd.tidy

## Card

Clean up and sync your `go.mod` and `go.sum` files.

```bash
go mod tidy
```

### Subcard: Remove Unused

Removes imports that are no longer referenced in your codebase.

### Subcard: Add Missing

Downloads and adds modules referenced in your codebase but missing from `go.mod`.

## Example

Full scenario of tidying, downloading, and cleaning module manifests after package changes:

1. Let's say you add a new third-party dependency import to `main.go`, and delete an old one:

`main.go`:
```go
package main

import (
	"fmt"
	"github.com/google/uuid" // Added new package
	// "github.com/fatih/color" <- Deleted old import
)

func main() {
	fmt.Println(uuid.New())
}
```

2. Run `go mod tidy` to clean up and sync files:
```bash
go mod tidy

# Expected Output:
# go: finding module for package github.com/google/uuid
# go: downloading github.com/google/uuid v1.6.0
```

3. Open `go.mod` to see the results:
```txt
module myapp

go 1.24

require github.com/google/uuid v1.6.0 // indirect
```
*(Notice that the old package `github.com/fatih/color` was completely removed, and the new `github.com/google/uuid` was automatically added!)*

## Deep

`go mod tidy` performs static analysis of your Go files to find all imports. It then ensures that `go.mod` matches the state of the imports in your source code:
- Adding dependencies that are imported but not listed.
- Removing dependencies that are listed but not imported.
- Syncing and tidying up `go.sum` (cryptographic hashes).

## Gotchas

- Always run `go mod tidy` before committing your code or building a container to ensure you have a clean and correct dependency manifest.

## Related

- setup.project
