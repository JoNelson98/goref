# slice.filter

## Card

Filter elements from a slice in-place (without allocating a new underlying array).

```go
n := 0
for _, x := range s {
	if keep(x) {
		s[n] = x
		n++
	}
}
s = s[:n]
```

### Subcard: Filter allocating a new slice

```go
var filtered []Type
for _, x := range s {
	if keep(x) {
		filtered = append(filtered, x)
	}
}
```

## Example

```go
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Filter in-place: Keep only even numbers
	n := 0
	for _, x := range numbers {
		if x%2 == 0 {
			numbers[n] = x
			n++
		}
	}
	// Shrink the slice to the number of kept elements
	numbers = numbers[:n]

	fmt.Println("Filtered evens (in-place):", numbers) // [2, 4, 6, 8, 10]
}
```

## Deep

The in-place filtering idiom is highly memory-efficient because it does not allocate any new memory on the heap.
- It reuses the same underlying array of the original slice.
- The index variable `n` tracks the write pointer. Since `n` is always less than or equal to the read index, we never overwrite any element before we read it.
- After the loop completes, we slice the original slice up to `n` (`s[:n]`) to shrink its length while retaining the same underlying capacity.

## Gotchas

- **Garbage Collection Warning:** If your slice contains pointers or objects containing pointers, the elements between `n` and `len(original)` are still referenced by the underlying array, preventing the garbage collector from freeing them! To avoid memory leaks, zero out the truncated elements before shrinking:
  ```go
  // Prevent memory leaks
  for i := n; i < len(s); i++ {
      s[i] = nil // or zero value
  }
  s = s[:n]
  ```

## Related

- slice.basics
- slice.remove-item
