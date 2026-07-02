# map.delete-key

## Card

Use the built-in `delete` function to remove a key from a map.

```go
delete(m, key)
```

### Subcard: Simple Delete

```go
delete(myMap, "key")
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

	// Delete key "apple"
	delete(m, "apple")

	fmt.Println(m) // map[banana:2]
}
```

## Deep

The `delete` function is built-in and does not return any value. If the key is not present in the map, or if the map is `nil`, `delete` is a no-op (it does nothing and does not panic).

## Gotchas

- Calling `delete` on a `nil` map is safe and won't panic, but *assigning* to a `nil` map will panic.
- `delete` does not release the map's underlying memory immediately. To release memory for very large maps, you may need to let the map go out of scope and be garbage collected, or copy surviving elements to a new map.

## Related

- map.check-exists
