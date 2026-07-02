# slice.basics

## Card

A slice is a dynamically-sized, flexible wrapper over an underlying array. Use `make`, `len`, and `append` to manage slices.

```go
s := []int{1, 2, 3}      // Literal initialization
s = append(s, 4)         // Append element (returns new slice)
l := len(s)              // Length of slice
```

### Subcard: Allocation with Make

```go
s := make([]int, 0, 10) // length 0, capacity 10 (highly optimized!)
```

## Example

```go
package main

import "fmt"

func main() {
	// 1. Initialize slice with make (capacity avoids subsequent allocations)
	items := make([]string, 0, 5)

	// 2. Append values
	items = append(items, "apple")
	items = append(items, "banana")

	fmt.Println("Items:", items)       // [apple banana]
	fmt.Println("Length:", len(items)) // 2

	// 3. Slice slicing (sub-slices share memory!)
	firstItem := items[0:1]
	fmt.Println("First:", firstItem) // [apple]
}
```

## Deep

Slices do not store any data themselves. They are a 3-word header containing:
1. A pointer to the underlying array.
2. The length (number of elements in the slice).
3. The capacity (maximum size the underlying array can grow to before a new array is allocated).

When you append past capacity, Go automatically allocates a new, larger underlying array (usually doubling capacity) and copies the old elements over. Using `make([]Type, len, cap)` allocates the exact memory size up front, avoiding this slow resize overhead.

## Gotchas

- **Slices Share Memory:** Sub-slicing (e.g. `sub := s[0:2]`) creates a new slice header pointing to the **same** underlying array! Modifying `sub[0]` will silently modify `s[0]`. To create an independent copy, allocate a new slice and use the built-in `copy` function:
  ```go
  dup := make([]int, len(s))
  copy(dup, s)
  ```
- **Appending returns new slice:** `append` can return a new slice pointing to a completely different array if capacity was exceeded. Always reassign back: `s = append(s, val)`.

## Related

- slice.remove-item
