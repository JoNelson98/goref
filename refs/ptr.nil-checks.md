# ptr.nil-checks

## Card

Always check if a pointer is `nil` before dereferencing it to avoid runtime panics.

```go
if p != nil {
    fmt.Println(*p)
}
```

### Subcard: Nil Return Checks

```go
val, err := fetchValue()
if err != nil {
    return nil, err
}
// Safe to use val if err is nil
```

## Example

```go
package main

import "fmt"

type Person struct {
	Name *string
}

func main() {
	p := Person{}

	// p.Name is nil because it's a pointer and wasn't initialized
	if p.Name != nil {
		fmt.Println(*p.Name)
	} else {
		fmt.Println("Name is not set (nil)")
	}

	// Initializing and assigning
	name := "Alice"
	p.Name = &name

	if p.Name != nil {
		fmt.Println("Name is:", *p.Name)
	}
}
```

## Deep

In Go, pointers default to `nil` (their zero value) when declared without an assignment. A `nil` pointer points to nothing.

Attempting to read or write to a `nil` pointer causes a **runtime panic: invalid memory address or nil pointer dereference**. This will crash your application. Rigorous nil-checking is an essential safety practice in production Go code.

## Gotchas

- **Methods on Nil Receivers:** Go allows calling methods on `nil` pointer receivers! If a method is defined on a pointer receiver `func (t *Type) Method()`, and you call it on a `nil` variable, the method runs. It will only panic if the method code itself attempts to dereference `t` without checking for `nil`.
  ```go
  func (p *Person) GetName() string {
      if p == nil {
          return "Guest"
      }
      return *p.Name // Still check if Name is nil!
  }
  ```

## Related

- ptr.basics
