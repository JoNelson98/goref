# cmd.build

## Card

Compile your Go application into a single, self-contained binary executable.

```bash
# Standard compile (defaults name to the directory name, e.g. "GOREF")
go build

# Compile with a custom name
go build -o myapp .

# Cross-compilation (Build for Linux from macOS)
GOOS=linux GOARCH=amd64 go build -o myapp-linux .
```

### Subcard: Shrink Binary Size

Strip debugging information and symbols to reduce file size by 30-50%:
```bash
go build -ldflags="-s -w" -o myapp .
```

## Example

A full build scenario for packaging a clean Go application with custom binary targets and optimizations:

1. Verify your Go program runs and compiles locally:
```bash
go run .
```

2. Compile without the `-o` flag (Go names the binary after the current directory):
```bash
# Suppose your current working folder is named "GOREF"
go build .

# Verify the newly created binary (named after the directory):
ls
# Output:
# GOREF    main.go   tui.go    go.mod
```

3. Compile with an explicit custom name inside a target directory:
```bash
go build -o build/myapp .
```

4. Compile an optimized, production-grade binary with stripped debug symbols (drastically reduces file size):
```bash
# -s strips debugging symbols, -w strips DWARF DWARF symbol tables
go build -ldflags="-s -w" -o build/myapp-prod .

# Check file sizes to see the savings:
ls -lh build/
# Output:
# -rwxr-xr-x  1 gopher  staff   6.4M Jun 30 12:00 myapp
# -rwxr-xr-x  1 gopher  staff   4.2M Jun 30 12:01 myapp-prod  (35% smaller!)
```

5. Cross-compile for a remote Linux server (AMD64 architecture) from your local machine:
```bash
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o build/myapp-linux-amd64 .
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

### Cross-Compilation
Cross-compilation is built-in. Set `GOOS` (operating system: `linux`, `darwin`, `windows`, `freebsd`) and `GOARCH` (architecture: `amd64`, `arm64`, `386`) variables before running the build command. No toolchains required!

## Gotchas

- **Do not compile single files:** Always run `go build .` or `go build -o name .` rather than `go build main.go` to ensure the Go compiler includes all package files in the directory.
- **Go binary includes runtime:** A Go binary is relatively large (usually 5MB-15MB) because it statically links the Go runtime (including the garbage collector and scheduler) so it can run on target machines without any Go installation.

## Related

- setup.project
- cmd.tidy
- cmd.run
