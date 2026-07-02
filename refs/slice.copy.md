# slice.copy

## Card

Create an independent clone of a slice that does not share memory with the original using the built-in `copy` function.

```go
dst := make([]Type, len(src))
copy(dst, src)
```

### Subcard: Simple Inline Copy

```go
dup := append([]int(nil), src...)
```

## Example

```go
package main

import "fmt"

func main() {
	original := []string{"A", "B", "C"}

	// 1. INCORRECT: Shallow assignment (shares memory)
	shallow := original
	shallow[0] = "Z"
	fmt.Println("Original after shallow write:", original) // [Z, B, C] (oops, mutated!)

	// Reset original
	original[0] = "A"

	// 2. CORRECT: Allocate and copy (independent clone)
	clone := make([]string, len(original))
	copiedElements := copy(clone, original)

	clone[0] = "Y"
	fmt.Println("Original after clone write:", original) // [A, B, C] (original is safe!)
	fmt.Println("Clone content:", clone)                 // [Y, B, C]
	fmt.Printf("Copied %d elements.\n", copiedElements)
}
```

## Deep

In Go, a slice is a header struct consisting of a pointer to an underlying array, a length, and a capacity. 
- When you do `shallow := original`, Go only copies the slice header (the pointer, length, and capacity). The pointer still points to the exact same underlying array. Writing to `shallow[i]` modifies the array shared by `original`.
- The built-in `copy(dst, src)` function copies elements from a source slice to a destination slice. It returns the number of elements copied, which is the **minimum** of `len(dst)` and `len(src)`.
- If the destination slice is not allocated with enough length (e.g. `dst := []int{}`), `copy` will copy **0 elements** and do nothing silently.

## Gotchas

- **Empty Destination Pitfall:** Always allocate `dst` with `len(src)` before calling `copy`. A common mistake is writing `var dst []int; copy(dst, src)` which copies nothing because `dst` has a length of 0.
- **Copying Overlapping Slices:** `copy` correctly handles overlapping slices within the same underlying array (e.g., `copy(s[1:], s[:3])`).

## Related

- slice.basics
- slice.filter
