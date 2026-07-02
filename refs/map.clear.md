# map.clear

## Card

Empty all entries from a map using the built-in `clear` function (Go 1.21+).

```go
clear(myMap)
```

### Subcard: Legacy Loop Clear (< Go 1.21)

```go
for k := range myMap {
    delete(myMap, k)
}
```

## Example

```go
package main

import "fmt"

func main() {
	m := map[string]int{
		"apple":  1,
		"banana": 2,
	}

	// Clear the map (keeps the allocated map header, zeroes elements)
	clear(m)

	fmt.Println("Map length after clear:", len(m)) // 0
	fmt.Println("Map is nil?", m == nil)            // false (map is still initialized!)
}
```

## Deep

The built-in `clear` function was introduced in Go 1.21.
- When applied to a map, `clear` removes all elements, leaving the map empty (`len(m) == 0`).
- Unlike assigning `m = nil` (which deallocates/uninitializes the reference, causing writes to panic), `clear` preserves the allocated map header, meaning the map remains ready for safe subsequent writes.
- Efficient: `clear` optimized loops internally or frees resources directly.

## Gotchas

- **Underlying Memory Allocation:** Calling `clear` or looping `delete` removes references so garbage collection can clean up values, but it **does not release the map's underlying bucket memory allocation** back to the operating system immediately. To completely reclaim large map memory, let the map go out of scope and re-initialize it:
  ```go
  m = make(map[string]int) // Re-allocate fresh map header
  ```

## Related

- map.delete-key
- map.check-exists
