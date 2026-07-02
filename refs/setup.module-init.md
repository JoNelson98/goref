# setup.module-init

## Card

Initialize a new Go module in the current directory.

```bash
go mod init <module-path>
```

### Subcard: Simple Init

```bash
go mod init github.com/user/project
```

## Example

```bash
mkdir myapi
cd myapi
go mod init github.com/user/myapi
```

This creates a `go.mod` file:
```txt
module github.com/user/myapi

go 1.24
```

## Deep

A Go module is a collection of Go packages stored in a file tree with a `go.mod` file at its root. The `go.mod` file defines:
- The module's path (which is also the import prefix for packages within the module).
- The minimum Go version required to build the module.
- The module's dependency requirements (with specific semantic versions).

## Gotchas

- **Do not use a generic name:** If you plan to publish your package, the module path must match the repository URL (e.g., `github.com/username/reponame`) so that `go get` can find and download it.
- **Avoid spaces or special characters:** Stick to alphanumeric characters, dashes, and dots.

## Related

- setup.project
- cmd.tidy
