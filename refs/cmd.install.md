# cmd.install

## Card

Build and install Go binaries globally on your system.

```bash
# Compile and install the executable in $GOPATH/bin (or $GOBIN)
go install .

# Install a specific remote package/tool globally without downloading source manually
go install github.com/user/tool@latest
```

### Subcard: Ensure Binaries are Executable Globally

Add Go's binary directory to your system shell configuration (`~/.bashrc`, `~/.zshrc`, or `~/.profile`):
```bash
export GOPATH=$(go env GOPATH)
export PATH=$PATH:$GOPATH/bin
```

---

## Example

### Scenario A: Standard Go Installation (`go install`)

1. Verify that your environment has `GOPATH/bin` in its search path:
```bash
go env GOPATH
# Output: /Users/username/go

echo $PATH | grep "$(go env GOPATH)/bin"
# If empty, append it in your ~/.zshrc or ~/.bashrc:
# export PATH="$PATH:$(go env GOPATH)/bin"
```

2. Run `go install` from the root of your Go project containing `package main`:
```bash
go install .
```

3. Run the compiled binary instantly from anywhere in your terminal:
```bash
myapp-binary-name --version
```

4. To update the global binary after changing local code or resources, simply run the command again:
```bash
go install .
```

---

### Scenario B: Manual Compile & Global Symlink (Ideal for Development)

Using a symlink allows you to rebuild the binary locally and have the changes instantly accessible globally without re-copying.

1. Compile the binary to a local path:
```bash
go build -o build/myapp .
```

2. Create a symbolic link in a directory that is already in your global `$PATH` (e.g., `/usr/local/bin`):
```bash
sudo ln -sf "$(pwd)/build/myapp" /usr/local/bin/myapp
```

3. To update the global command, simply rebuild the local binary. The symlink automatically points to the fresh build:
```bash
go build -o build/myapp .
```

---

## Deep

- **`go install` compilation:** When you run `go install`, Go compiles your package and any dependencies, and then moves the resulting executable binary to the directory named by the `GOBIN` environment variable. If `GOBIN` is not set, it defaults to the `bin` subdirectory of the first directory listed in the `GOPATH` environment variable (typically `~/go/bin`).
- **Caching:** `go install` is highly efficient. It leverages the Go build cache, compiling only what has changed since the last build.
- **Symlinking vs. Copying:** Creating a symlink (`ln -sf`) makes local development iteration incredibly fast because the active global command directly references your build workspace output. Copying (`cp`) isolates the global binary from your workspace directory, which is safer for finished production utilities but requires manually copying the executable again on every update.

---

## Gotchas

- **Non-main package installation:** If you run `go install` inside a library package (i.e., a package other than `package main`), Go compiles the package to verify it builds correctly and caches the package object, but **does not** generate a global executable.
- **Go version mismatches:** When installing packages using `@latest` from remote repositories, ensure your local Go compiler version meets the minimum version required by the module's `go.mod` file.
- **Elevated permissions:** Creating symlinks or copying binaries to global directories like `/usr/local/bin` or `/usr/bin` typically requires `sudo` privileges.

---

## Related

- cmd.build
- setup.project
