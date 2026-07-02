# ptr.basics

## Card

A pointer stores the memory address of a value. Use `&` to get the address, and `*` to dereference (read/write the value).

```go
p := &val  // p is a pointer to val
v := *p    // v is the value p points to
```

### Subcard: Declaring Pointers

```go
var p *int // p is a nil pointer to an int
```

## Example

```go
package main

import "fmt"

func main() {
	x := 42
	p := &x // Get memory address of x

	fmt.Println(p)  // Prints address, e.g., 0xc0000140a8
	fmt.Println(*p) // Prints 42 (dereferencing)

	*p = 21        // Change value at address p
	fmt.Println(x) // Prints 21 (x is updated!)
}
```

## Deep

Go is a "pass-by-value" language. When you pass an argument to a function, Go copies the value. If you pass a struct or primitive, the function cannot modify the original.

Passing a pointer (`*Type`) copies the memory address. This allows the function to modify the value in the caller's scope, and is also more efficient for very large structs as it avoids copying their entire data.

## Gotchas

- **Nil Dereference:** Dereferencing a `nil` pointer (e.g., `var p *int; *p = 5`) will panic immediately. Always check for `nil` or allocate memory before writing!
- **Pointers to Slices/Maps:** You almost never need pointers to slices (`*[]type`) or maps (`*map[k]v`) because slices and maps are already header structures that reference underlying data. Passing them directly allows modification of their contents.

## Related

- ptr.nil-checks
- map.check-exists
