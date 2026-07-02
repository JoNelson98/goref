# struct.basics

## Card

A struct is a typed collection of fields. Use struct fields with backtick tags to map Go fields to JSON properties.

```go
type Todo struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
}
```

### Subcard: Instantiate Struct

```go
// Positional
t1 := Todo{1, "Buy milk"}

// Named Fields (Highly recommended!)
t2 := Todo{
    ID:    2,
    Title: "Clean room",
}
```

## Example

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
}

func main() {
	u := User{
		ID:       101,
		Username: "gopher",
		IsAdmin:  true,
	}

	// Marshalling (Struct -> JSON bytes)
	bytes, _ := json.Marshal(u)
	fmt.Println(string(bytes)) // {"id":101,"username":"gopher","is_admin":true}
}
```

## Deep

Structs in Go are custom user-defined types. They are highly efficient because fields are laid out contiguously in memory.
- **Exporting Fields:** Just like package functions, struct fields must start with an **uppercase letter** to be exported (visible outside the package). If a field starts with a lowercase letter, standard JSON encoders (`encoding/json`) cannot see or marshal/unmarshal it!
- **Struct Tags:** Backtick tags like `` `json:"id"` `` tell the `encoding/json` package exactly how to serialize the field. If you omit the tag, the JSON key defaults to the uppercase field name (e.g., `ID` or `Title`).

## Gotchas

- **Lowercase Fields are Hidden:** If you declare a field as `id int` (lowercase), `json.Marshal` will silently ignore it, and `json.Unmarshal` will never set its value from JSON payloads. Always use `ID int` (uppercase).
- **Nil struct fields:** Pointers inside structs are initialized to `nil`.

## Related

- pkg.basics
- json.decode-request
