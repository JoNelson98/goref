# slice.remove-item

## Card

Remove an item from a slice by slicing and appending.

```go
// Preserve order
s = append(s[:i], s[i+1:]...)

// Fast (does NOT preserve order)
s[i] = s[len(s)-1]
s = s[:len(s)-1]
```

### Subcard: Keep Order

```go
s := []int{1, 2, 3, 4}
i := 1 // index to remove
s = append(s[:i], s[i+1:]...) // [1, 3, 4]
```

### Subcard: Fast Delete (Unordered)

```go
s := []int{1, 2, 3, 4}
i := 1 // index to remove
s[i] = s[len(s)-1] // Copy last element to index i -> [1, 4, 3, 4]
s = s[:len(s)-1]   // Slice off last element -> [1, 4, 3]
```

## Example

```go
package main

import "fmt"

func main() {
	// 1. Order-preserving remove
	ordered := []string{"A", "B", "C", "D"}
	idx := 1 // Remove "B"
	ordered = append(ordered[:idx], ordered[idx+1:]...)
	fmt.Println("Ordered:", ordered) // [A, C, D]

	// 2. Unordered fast remove
	fast := []string{"A", "B", "C", "D"}
	idx = 1 // Remove "B"
	fast[idx] = fast[len(fast)-1]
	fast = fast[:len(fast)-1]
	fmt.Println("Fast Unordered:", fast) // [A, D, C]
}
```

## Deep

Since slices are backed by arrays, deleting an element from the middle of a slice requires shifting all subsequent elements to the left, which is an $O(N)$ operation. 

The unordered approach is $O(1)$ because it only replaces the targeted element with the last element and shrinks the slice.

## Gotchas

- **Memory Leak Warning:** If the slice contains pointers or structs with pointers, removing an element using `s = append(s[:i], s[i+1:]...)` or the fast swap method leaves the duplicate element at the end of the underlying array allocated. It is not garbage collected because the underlying array still holds a reference! To avoid memory leaks, set the deleted index/slot to `nil` (or zero-value) before shrinking:
  ```go
  // For pointers
  copy(s[i:], s[i+1:])
  s[len(s)-1] = nil // Avoid memory leak
  s = s[:len(s)-1]
  ```

## Related

- map.delete-key
