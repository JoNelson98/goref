# struct.methods

## Card

Define behavior for structs using value receivers or pointer receivers.

```go
// Value receiver (reads data, works on a copy)
func (t MyStruct) View() string {}

// Pointer receiver (writes/mutates, works on original)
func (t *MyStruct) Increment() {}
```

### Subcard: Pointer Receiver Syntax

```go
func (p *Person) SetName(name string) {
	p.Name = name // Mutates original struct
}
```

## Example

```go
package main

import "fmt"

type Counter struct {
	Value int
}

// Value Receiver: Mutates a COPY, original unchanged
func (c Counter) AddValue(val int) {
	c.Value += val
}

// Pointer Receiver: Mutates the ORIGINAL struct
func (c *Counter) AddPointer(val int) {
	c.Value += val
}

func main() {
	c := Counter{Value: 10}

	c.AddValue(5)
	fmt.Println("After Value:", c.Value) // Still 10!

	c.AddPointer(5)
	fmt.Println("After Pointer:", c.Value) // 15!
}
```

## Deep

Methods in Go are just functions with a special "receiver" argument between the `func` keyword and the method name.
- **Value Receiver (`T`)**: Go passes a copy of the struct to the method. Any modifications are made to the copy and discarded when the method returns.
- **Pointer Receiver (`*T`)**: Go passes the memory address of the struct. This allows direct modifications to the caller's struct fields, and avoids copying the struct on every call (highly efficient for larger structs).

## Gotchas

- **Consistency Rule:** If any method of your struct requires a pointer receiver, **all** methods of that struct should have pointer receivers to keep the interface consistent.
- **Automatic Dereferencing:** You don't need to manually dereference a pointer receiver to access fields (e.g. write `(*c).Value`), Go automatically dereferences `c.Value` for you.

## Related

- struct.basics
- ptr.basics
- ptr.flow-chart
