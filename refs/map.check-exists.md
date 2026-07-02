# map.check-exists

## Card

Use the "comma ok" idiom to check if a key exists in a map.

```go
val, ok := m[key]
if !ok {
    // key does not exist
}
```

### Subcard: Just Check Existence

If you don't need the value, use a blank identifier:
```go
if _, ok := m[key]; ok {
    // key exists
}
```

## Example

```go
package main

import "fmt"

func main() {
	m := map[string]int{"apple": 1}

	if val, ok := m["apple"]; ok {
		fmt.Printf("apple exists: %d\n", val)
	} else {
		fmt.Println("apple does not exist")
	}

	if _, ok := m["banana"]; !ok {
		fmt.Println("banana does not exist")
	}
}
```

## Deep

If you look up a non-existent key in a map without the "comma ok" idiom (e.g., `val := m[key]`), Go returns the zero value for the map's value type (e.g., `0` for `int`, `""` for `string`, `nil` for pointers). You cannot distinguish between "key exists with a zero value" and "key does not exist" without using `ok`.

## Gotchas

- Looking up a key in a `nil` map is safe and returns the zero value, but writing to a `nil` map will panic. Always initialize maps before writing.

## Related

- map.delete-key
