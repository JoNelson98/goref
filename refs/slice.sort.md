# slice.sort

## Card

Sort a slice using the standard library `slices` package (Go 1.21+).

```go
import "slices"

// Sort primitives (ints, floats, strings)
slices.Sort(mySlice)

// Sort custom structs by custom fields
slices.SortFunc(users, func(a, b User) int {
	return cmp.Compare(a.Age, b.Age)
})
```

### Subcard: Legacy Sort Package (< Go 1.21)

```go
import "sort"

sort.Ints(myIntSlice)
sort.Slice(users, func(i, j int) bool {
	return users[i].Age < users[j].Age
})
```

## Example

```go
package main

import (
	"cmp"
	"fmt"
	"slices"
)

type User struct {
	Name string
	Age  int
}

func main() {
	// 1. Sort strings
	names := []string{"Charlie", "Alice", "Bob"}
	slices.Sort(names)
	fmt.Println("Sorted names:", names) // [Alice, Bob, Charlie]

	// 2. Sort structs by Age
	users := []User{
		{"Charlie", 35},
		{"Alice", 25},
		{"Bob", 30},
	}
	slices.SortFunc(users, func(a, b User) int {
		return cmp.Compare(a.Age, b.Age)
	})

	fmt.Printf("Sorted users: %+v\n", users)
}
```

## Deep

Go 1.21 introduced the generic `slices` package which is safer and significantly faster than the legacy `sort` package.
- `slices.Sort` works on any type that is ordered (type constraints supporting `<`, `>`, etc.).
- `slices.SortFunc` allows custom sorting logic using a comparator function that returns negative (less than), zero (equal), or positive (greater than).
- The `cmp.Compare` helper from the built-in `cmp` package is the standard, type-safe way to implement this comparator function.

## Gotchas

- **In-place Mutation:** Sorting is performed **in-place**. It modifies the original slice order. If you need to keep the original order, make a copy of the slice using `slice.copy` first before sorting.
- **Float Sorting:** Floating-point slices containing `NaN` values require care, as `NaN` comparison returns false. `slices.Sort` handles `NaN` correctly by placing them at the beginning.

## Related

- slice.basics
- slice.copy
