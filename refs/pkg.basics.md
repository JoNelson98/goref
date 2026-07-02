# pkg.basics

## Card

Go programs are organized into packages. A package is a directory containing one or more Go files.

```go
package mypkg

// Exported (Visible outside the package)
func Run() {}

// Unexported (Only visible inside the package)
func helper() {}
```

### Subcard: Capitalization Rule

- **Exported identifiers** (functions, types, variables, struct fields) start with an **uppercase letter** (e.g., `fmt.Println`).
- **Unexported identifiers** start with a **lowercase letter** (e.g., `helper()`) and are only visible within the same package.

## Example

```go
package mathutils

// Add is exported because it starts with an uppercase 'A'
func Add(a, b int) int {
	return a + b
}

// multiply is unexported because it starts with a lowercase 'm'
func multiply(a, b int) int {
	return a * b
}
```

Using the package:
```go
package main

import (
	"fmt"
	"mymodule/mathutils"
)

func main() {
	fmt.Println(mathutils.Add(2, 3)) // Works
	// mathutils.multiply(2, 3)      // Compile Error: unexported
}
```

## Deep

Every Go file must start with a `package` declaration. Packages serve as namespaces:
- They prevent naming collisions between different modules.
- They enforce clean API boundaries via capitalization (exported vs unexported).
- Importing packages is done via their fully-qualified import path starting with the module name defined in `go.mod`.

## Gotchas

- **Circular Dependencies:** Package `A` cannot import package `B` if package `B` also imports package `A`. Go strictly prohibits circular imports and will fail to compile. Keep packages orthogonal and decoupled.
- **Folder Names:** Keep package names identical to the directory name they live in. Avoid plural folder names (e.g., use `pkg/api` instead of `pkg/apis`).

## Related

- setup.flat-layout
- setup.project
