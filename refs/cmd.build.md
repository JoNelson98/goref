# cmd.build

## Card

Compile your Go application into a single, self-contained binary executable.

```bash
# Standard compile
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

2. Compile a standard binary for your local machine's operating system:
```bash
go build -o build/myapp .
```

3. Compile an optimized, production-grade binary with stripped debug symbols (drastically reduces file size):
```bash
# -s strips debugging symbols, -w strips DWARF symbol tables
go build -ldflags="-s -w" -o build/myapp-prod .

# Check file sizes to see the savings:
ls -lh build/
# Output:
# -rwxr-xr-x  1 gopher  staff   6.4M Jun 30 12:00 myapp
# -rwxr-xr-x  1 gopher  staff   4.2M Jun 30 12:01 myapp-prod  (35% smaller!)
```

4. Cross-compile for a remote Linux server (AMD64 architecture) from your local machine:
```bash
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o build/myapp-linux-amd64 .
```

## Deep

The `go build` command compiles the packages along with their dependencies.
- If you build a `package main`, it generates an executable binary.
- If you build any other package, it compiles the files to verify they build clean but discards the output (caches it for subsequent builds).
- Cross-compilation is built-in. Set `GOOS` (operating system: `linux`, `darwin`, `windows`, `freebsd`) and `GOARCH` (architecture: `amd64`, `arm64`, `386`) variables before running the build command. No toolchains required!

## Gotchas

- **Do not compile single files:** Always run `go build .` or `go build -o name .` rather than `go build main.go` to ensure the Go compiler includes all package files in the directory.
- **Go binary includes runtime:** A Go binary is relatively large (usually 5MB-15MB) because it statically links the Go runtime (including the garbage collector and scheduler) so it can run on target machines without any Go installation.

## Related

- setup.project
- cmd.tidy
- cmd.run
