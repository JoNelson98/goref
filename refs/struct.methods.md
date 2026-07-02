# struct.methods

## Card

Define behavior for custom types using value receivers or pointer receivers.

```go
// Value receiver: Operates on a COPY. Safe, read-only, thread-safe.
func (c Counter) Read() int { return c.Value }

// Pointer receiver: Operates on the ORIGINAL. Mutates state, high performance.
func (c *Counter) Increment() { c.Value++ }
```

### Subcard: Method Automatic Referencing

Go automatically handles conversions between values and pointers when calling methods:
```go
c := Counter{Value: 10}
c.Increment()  // Go automatically converts this to: (&c).Increment()

cp := &Counter{Value: 10}
_ = cp.Read()  // Go automatically converts this to: (*cp).Read()
```

## Example

Comprehensive example demonstrating mutations, stack copy behaviors, and compiler limits on addressability:

```go
package main

import "fmt"

type Point struct {
	X, Y int
}

// 1. Value Receiver: Receives a Stack Copy
func (p Point) ScaleValue(factor int) {
	p.X *= factor
	p.Y *= factor
	fmt.Printf("Inside Value method (Copy): {%d, %d}\n", p.X, p.Y)
}

// 2. Pointer Receiver: Receives the original memory address
func (p *Point) ScalePointer(factor int) {
	p.X *= factor
	p.Y *= factor
	fmt.Printf("Inside Pointer method (Original): {%d, %d}\n", p.X, p.Y)
}

func main() {
	p := Point{X: 2, Y: 3}

	// Call Value Method
	p.ScaleValue(10)                       // Inside: {20, 30}
	fmt.Printf("In Main after Value call: {%d, %d}\n", p.X, p.Y) // Still {2, 3}!

	// Call Pointer Method
	p.ScalePointer(10)                       // Inside: {20, 30}
	fmt.Printf("In Main after Pointer call: {%d, %d}\n\n", p.X, p.Y) // Mutated to {20, 30}!

	// --- CRUCIAL ADDRESSABILITY LIMITATION GOTCHA ---
	pointsMap := map[string]Point{
		"start": {X: 1, Y: 1},
	}
	
	// Map values are NOT addressable! Go cannot take the address of a map element.
	// This will throw a compile error:
	// pointsMap["start"].ScalePointer(2) 
	
	_ = pointsMap
}
```

## Deep

Understanding the difference between value and pointer receivers is fundamental to writing correct and efficient Go code.

### Value Receiver (`T`) Mechanics:
- **Stack Allocation:** Go copies the entire struct value onto the method's local call stack.
- **Immutability:** Modifications are isolated to the local copy. It is impossible to mutate the caller's struct.
- **Thread-Safety:** Reading from a value receiver is naturally thread-safe for that method's scope because the method operates strictly on its own private stack copy, meaning no concurrent goroutine can write to it.

### Pointer Receiver (`*T`) Mechanics:
- **Address Passing:** Go copies only the memory address of the struct (exactly 8 bytes on a 64-bit architecture) to the method call.
- **Mutation:** The method dereferences the address to mutate the original fields directly.
- **Zero Copying:** Unbelievably efficient for very large structs since we avoid the expensive overhead of copying multiple struct fields in memory on every call.

---

## When to Use: Comprehensive Checklist

### Use POINTER Receivers (`*T`) if:
1. **Mutation:** The method needs to modify fields in the receiver.
2. **Safety (Mutexes/Locks):** The struct contains synchronizing fields (like `sync.Mutex` or `sync.WaitGroup`) or resources (like network sockets, database connections, files). **These must never be copied!** If they are copied, their lock-states are copied, leading to deadlocks or memory corruption.
3. **Performance:** The struct is "large" (usually containing more than 5 fields, or arrays, or other heavy structs). Passing a pointer only copies 8 bytes of address space.
4. **Consistency:** If any other method on that struct type uses a pointer receiver, **all methods of that struct should use a pointer receiver** (even if they are read-only). Mixing them leads to severe confusion.

### Use VALUE Receivers (`T`) if:
1. **Small & Read-Only:** The struct is small, simple, and read-only (e.g., simple coordinate pairs `Point{X,Y}`, or currency `Money{Amount}`).
2. **Immutable Guarantee:** You want to guarantee that a method call can never mutate the original object under any circumstances.
3. **Reference Types:** The receiver is a reference type like a map, slice, channel, or function. Since these types are already internally lightweight header structs pointing to external data, passing them as values is cheap, and they can still modify the underlying elements directly!

## Gotchas

- **Addressability Errors:** You can only call pointer receivers on addressable values. Variables (`p`) and pointer variables (`&p`) are addressable. Values returned directly from maps (`m[key]`), constants, or literals (like `Point{1,2}.ScalePointer()`) are NOT addressable, and trying to call pointer methods on them will throw compile errors!
- **Silent No-Ops:** Writing to a field inside a value receiver method compiles perfectly but is a silent no-op relative to the caller. Always double-check your method signatures.

## Related

- struct.basics
- ptr.basics
- ptr.flow-chart
- ptr.pkg-usage
