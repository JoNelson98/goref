# cmd.run

## Card

Compile and run your Go program on the fly (useful for fast local feedback).

### Syntax
```bash
go run <path_to_package_or_files> [arguments...]
```

### Quick Example
```bash
# Compile and run the current folder program
go run .

# Compile and pass flags/arguments to your program
go run . -name=Alice
```

### Subcard: Run with Arguments

```bash
go run . list setup
```

## Example

Detailed scenario showing the use of `go run` during development to test multi-file main packages with command line flags:

1. Suppose we have a flat project layout inside the `myapp/` directory:

`main.go`:
```go
package main

import (
	"flag"
)

func main() {
	nameFlag := flag.String("name", "Gopher", "Name to greet")
	flag.Parse()
	greet(*nameFlag)
}
```

`greetings.go`:
```go
package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
```

2. Run the multi-file package using `go run .` and pass command-line arguments/flags:
```bash
# Run the entire local package and pass flags
go run . -name=Alice

# Expected Output:
# Hello, Alice!
```

3. **DO NOT DO THIS** (Running a single file will fail because it misses helper files):
```bash
go run main.go -name=Alice

# Expected Compile Error:
# # command-line-arguments
# ./main.go:10:2: undefined: greet
```

## Deep

`go run` compiles the files, creates a temporary executable binary in your system's temp directory, runs it, and then automatically cleans up and deletes the executable after termination. It is meant exclusively for fast local feedback during development, not for production.

## Gotchas

- **Never use `go run` in production:** It incurs compilation overhead on every launch and requires the Go SDK to be installed on the host machine. Statically compile executables using `go build` for deployment.
- **Do not import your own root:** Since all files in the directory share the same package, they do not need to (and cannot) import each other.

## Related

- setup.project
- cmd.build
- setup.flat-layout
