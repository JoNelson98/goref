# cmd.add-package

## Card

Add, download, or update external package dependencies in your Go project using `go get`.

### Syntax
```bash
go get <package_path>[@version]
```

### Quick Examples
```bash
# Add the latest version of UUID package
go get github.com/google/uuid

# Pin to a specific semantic version
go get github.com/google/uuid@v1.6.0

# Completely remove a dependency package
go get github.com/google/uuid@none
```

### Subcard: Update Dependencies

```bash
# Update a package to its latest minor or patch release
go get -u github.com/google/uuid
```

## Example

Detailed step-by-step walkthrough of installing and using an external UUID package in a Go module:

1. Create a minimal source file `main.go` importing the package:
```go
package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Printf("Generated UUID: %s\n", id.String())
}
```

2. Download the package and update `go.mod` using `go get`:
```bash
# Download and add uuid package to go.mod
go get github.com/google/uuid

# Expected Terminal Output:
# go: downloading github.com/google/uuid v1.6.0
# go: added github.com/google/uuid v1.6.0
```

3. Sync and verify your `go.sum` file:
```bash
go mod tidy
```

4. Run the code:
```bash
go run .
# Expected Output:
# Generated UUID: f47ac10b-58cc-4372-a567-0e02b2c3d479
```

## Deep

The `go get` command interacts directly with Go module proxies (like `proxy.golang.org`) to fetch dependencies:
- It downloads the requested module and its dependencies to your local Go module cache (`$GOPATH/pkg/mod`).
- It adds a line representing the module with its cryptographic checksum version to your project's `go.mod` file.
- It updates the cryptographic database inside your project's `go.sum` to prevent middle-man tampering during subsequent builds.

## Gotchas

- **Do not run outside of a module:** You cannot run `go get` unless you are inside a directory containing a valid `go.mod` file (or a subdirectory). Go will throw an error: `go: go.mod file not found in current directory or any parent directory`.
- **Remove via @none:** To completely remove a dependency from your caches and project files, run `go get path@none` and then clean up with `go mod tidy`.

## Related

- cmd.tidy
- setup.project
- cmd.build
