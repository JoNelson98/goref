# setup.flat-layout

## Card

A flat project layout keeps all Go files in the root directory under the same package (usually `package main`).

```txt
myapp/
  go.mod
  main.go
  tui.go
  cards.go
```

### Subcard: Run Flat Project

```bash
go run .
```

Do not run `go run main.go` as it ignores other package files.

## Example

An example of a flat layout where `main.go` and `tui.go` share variables and functions directly because they are in the same package:

`main.go`:
```go
package main

import "fmt"

func main() {
	runApp()
}
```

`tui.go`:
```go
package main

import "fmt"

func runApp() {
	fmt.Println("Running TUI...")
}
```

## Deep

For small to medium-sized applications, libraries, and CLIs, keeping all Go files in the root directory (`package main` or your library package) is the idiomatic, simple Go way. You don't need complex subdirectories (like `/src/`, `/pkg/`, or `/internal/`) until your project grows large.

All files in a flat directory belong to the *same* package, meaning they share all types, variables, and functions globally within that folder, even if they are unexported (lowercase).

## Gotchas

- **Single Package per Folder:** A single directory can only contain files belonging to exactly **one** Go package. You cannot have files with `package main` and `package auth` in the same directory.
- **Do not import your own root:** Since all files in the directory share the same package, they do not need to (and cannot) import each other.

## Related

- setup.project
- pkg.basics
