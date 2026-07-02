# slice.contains

## Card

Check if a slice contains a specific value using `slices.Contains`.

```go
import "slices"

exists := slices.Contains(mySlice, value)
```

### Subcard: Custom Field Match (Structs)

```go
exists := slices.ContainsFunc(users, func(u User) bool {
	return u.Name == "Alice"
})
```

## Example

```go
package main

import (
	"fmt"
	"slices"
)

type Product struct {
	ID   int
	Name string
}

func main() {
	fruits := []string{"apple", "banana", "orange"}

	// 1. Primitives check
	if slices.Contains(fruits, "banana") {
		fmt.Println("Banana exists!")
	}

	// 2. Structs check by field
	products := []Product{
		{1, "Laptop"},
		{2, "Phone"},
	}

	hasPhone := slices.ContainsFunc(products, func(p Product) bool {
		return p.Name == "Phone"
	})
	fmt.Println("Has Phone in inventory:", hasPhone) // true
}
```

## Deep

Before Go 1.21, checking for element existence required writing custom `for` loops or building map lookup caches. 
- The generic helper `slices.Contains` runs a simple linear scan $O(N)$ lookup. It is perfect and efficient for small to medium-sized slices.
- If you have huge datasets and do frequent membership lookups, do not scan slices! Convert the slice to a map (`map[Type]struct{}`) which provides $O(1)$ constant-time lookup.

## Gotchas

- **Linear Scan Performance:** `slices.Contains` is $O(N)$. Using it repeatedly inside loops can lead to $O(N^2)$ quadratic slowdowns. Use maps for fast lookup tables.

## Related

- map.check-exists
- slice.basics
