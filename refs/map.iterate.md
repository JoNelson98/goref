# map.iterate

## Card

Iterate over all keys and values in a map using a `for range` loop.

```go
for key, value := range myMap {
    fmt.Println(key, value)
}
```

### Subcard: Iterate in Sorted Order

```go
keys := make([]KeyType, 0, len(myMap))
for k := range myMap {
    keys = append(keys, k)
}
slices.Sort(keys) // Sort keys alphabetically/numerically
for _, k := range keys {
    fmt.Println(k, myMap[k])
}
```

## Example

```go
package main

import (
	"fmt"
	"slices"
)

func main() {
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	// 1. Non-deterministic iteration
	fmt.Println("--- Standard Loop ---")
	for k, v := range m {
		fmt.Printf("%s: %d\n", k, v)
	}

	// 2. Sorted iteration (consistent output)
	fmt.Println("--- Sorted Loop ---")
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, m[k])
	}
}
```

## Deep

Map iteration in Go is **non-deterministic**. 
- Go deliberately randomizes the starting index of map iteration every time a `for range` is evaluated. This is a deliberate language feature to prevent developers from relying on any implicit insertion order.
- To iterate consistently or in a specific order, you must extract the map keys to a slice, sort the slice, and loop over the sorted slice of keys to pull values from the map.

## Gotchas

- **Never rely on map order:** Tests that depend on raw map iteration will flake/fail intermittently because map ordering changes across executions and different Go runtime compiler versions. Always sort keys for test assertions.
- **Modifying map during iteration:** It is completely safe to delete or add keys to a map during iteration in the same goroutine. However, adding keys during iteration is non-deterministic: the loop may or may not visit the newly added elements.

## Related

- map.check-exists
- slice.sort
