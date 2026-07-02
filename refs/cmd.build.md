# cmd.build

## Card

Compile your Go application into a single, self-contained binary executable.

### Syntax
```bash
go build -o <output_name> <path_to_package>
```

### Quick Example
```bash
# Compile the current folder package into a binary named "myapp"
go build -o myapp .

# Cross-compilation (Build for Linux AMD64 from macOS)
GOOS=linux GOARCH=amd64 go build -o myapp-linux .
```

### Subcard: Shrink Binary Size

Strip debugging information and symbols to reduce file size by 30-50%:
```bash
go build -ldflags="-s -w" -o myapp .
```

## Example

Complete walkthrough of creating, compiling, and running a standalone Go binary:

1. Create a minimal source file `main.go`:
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, production-ready GoRef!")
}
```

2. Compile your application into a custom-named executable binary:
```bash
# Compile current package and output executable to build/ folder
go build -o build/myapp .
```

3. Run the compiled standalone binary (no Go SDK required to run!):
```bash
./build/myapp

# Expected Output:
# Hello, production-ready GoRef!
```

4. Optimize and shrink the binary size for production release by stripping debug symbols:
```bash
go build -ldflags="-s -w" -o build/myapp-prod .

# Verify the file size savings:
ls -lh build/
# Expected Output:
# -rwxr-xr-x  1 gopher  staff   6.4M Jun 30 12:00 myapp
# -rwxr-xr-x  1 gopher  staff   4.2M Jun 30 12:01 myapp-prod  (35% smaller!)
```

5. Cross-compile the exact same code for a Linux server (AMD64 architecture) from your local machine:
```bash
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o build/myapp-linux .
```

## Deep

### Default Binary Naming Rules
If you execute a build command **without the `-o` flag**, Go determines the output executable binary name based on how you target your compilation:

1. **Building the Directory (`go build` or `go build .`)**: The compiler names the executable exactly after the **parent directory name**. 
   * *Example:* Running `go build` inside `/Users/jn/Documents/GOREF` produces a binary named **`GOREF`** (or `goref` depending on filesystem case conventions).
2. **Building a Subdirectory Path (`go build ./cmd/mytool`)**: The compiler names the executable after the **last folder segment** of the path.
   * *Example:* Running `go build ./cmd/app-server` produces a binary named **`app-server`**.
3. **Building Specific Files (`go build main.go`)**: The compiler names the executable after the **first file argument** minus the `.go` extension.
   * *Example:* Running `go build main.go` produces a binary named **`main`**.

### Cross-Compilation Mechanics
Go's compilation is unique because the compiler and linker are written entirely in Go and support cross-compilation out of the box with zero external dependencies, linkers, or GCC toolchains. Simply set the target `GOOS` and `GOARCH` environment variables, and the compiler compiles a native binary for that operating system.

## Gotchas

- **Do not compile single files:** Always run `go build .` or `go build -o name .` rather than `go build main.go` to ensure the Go compiler includes all package files in the directory.
- **Go binary includes runtime:** A Go binary is relatively large (usually 5MB-15MB) because it statically links the Go runtime (including the garbage collector and scheduler) so it can run on target machines without any Go installation.

## Related

- setup.project
- cmd.tidy
- cmd.run
