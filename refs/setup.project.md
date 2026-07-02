# setup.project

## Card

Set up a new Go project from zero.

```bash
mkdir myapp
cd myapp
go mod init github.com/you/myapp
touch main.go
go run .
```

### Subcard: Create Folder

```bash
mkdir myapp
cd myapp
```

### Subcard: Init Module

```bash
go mod init github.com/you/myapp
```

### Subcard: Create Main File

```go
package main

import "fmt"

func main() {
	fmt.Println("hello")
}
```

## Example

A complete working example of a minimal Go program:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, GoRef!")
}
```

Run it with:
```bash
go run main.go
```

## Deep

A Go project starts with `go.mod` which defines the module path and Go version.

Start flat:
```txt
myapp/
  go.mod
  main.go
```

Only add folders (like `internal/`, `pkg/`, `cmd/`) when the project starts getting complex. For small CLI utilities or services, keeping everything in the root package is preferred and idiomatic.

## Gotchas

- Do not run `go mod init` in your home folder or a parent directory containing other Go projects.
- Prefer `go run .` over `go run main.go` when you have multiple files in the main package, as running just `main.go` will ignore other package files and cause build errors.

## Related

- cmd.run
- cmd.tidy
- setup.module-init
