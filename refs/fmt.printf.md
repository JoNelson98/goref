# fmt.printf

## Card

Format and print text to standard output using the `fmt.Printf` function and type-specific formatting "verbs" (placeholders).

```go
fmt.Printf("Hello %s, you are %d years old\n", "Alice", 30)
```

### Essential Formatting Verbs Cheat-Sheet

| Verb | Type / Description | Output Example |
| :--- | :--- | :--- |
| **`%v`** | **Default/Generic Value:** Best for rapid debug, prints any type | `3.14`, `true`, `map[a:1]` |
| **`%+v`** | **Detailed Struct:** Prints struct fields with their field names | `{Name:Alice Age:30}` |
| **`%#v`** | **Go-Syntax representation** of the value | `main.User{Name:"Alice"}` |
| **`%T`** | **Type:** Prints the Go type name of the value | `int`, `string`, `*main.User` |
| **`%s`** | **String** / byte slices | `"Hello, GoRef!"` |
| **`%d`** | **Integer** (base 10) | `42`, `-105` |
| **`%f`** | **Float** (decimal point) | `123.456000` (defaults to 6 decimals) |
| **`%.2f`**| **Float with custom precision** (rounds to 2 decimals) | `123.46` |
| **`%t`** | **Boolean** | `true` or `false` |
| **`%p`** | **Pointer Address** (base 16 hex notation with leading `0x`)| `0xc0000140a8` |
| **`%%`** | **Literal percent sign** | `%` |

## Example

Complete Go program showcasing formatting verbs, custom padding, and struct layouts:

```go
package main

import "fmt"

type Profile struct {
	Username string
	Admin    bool
}

func main() {
	user := Profile{"Gopher", true}

	// 1. Strings and Numbers
	fmt.Printf("Name: %s | Age: %d | Score: %.1f%%\n", "Alice", 28, 98.45)
	// Output: Name: Alice | Age: 28 | Score: 98.5%

	// 2. Generic Values and Types
	fmt.Printf("Default value: %v | Type: %T\n", user, user)
	// Output: Default value: {Gopher true} | Type: main.Profile

	// 3. Rich Struct Formatting (Highly recommended for debugging!)
	fmt.Printf("Struct fields: %+v\n", user)
	// Output: Struct fields: {Username:Gopher Admin:true}

	// 4. Pointer Formatting
	p := &user
	fmt.Printf("Pointer address: %p\n", p)
	// Output: Pointer address: 0xc0000a6040
}
```

## Deep

### Printing Functions Comparison
The standard library `fmt` package has three primary output functions:
1. **`fmt.Print(a...)`**: Concatenates arguments using their default representations. It does **not** add spaces between non-string arguments automatically, and does **not** append a newline.
2. **`fmt.Println(a...)`**: Prints arguments, adds spaces between each argument automatically, and **always appends a newline (`\n`)** at the end.
3. **`fmt.Printf(format, a...)`**: Compiles custom format templates. It **does not append a newline**; you must manually add `\n` to end the line.

### Under the Hood Performance Cost
`fmt.Printf` is extremely powerful but has a hidden runtime performance cost. It uses **runtime reflection** (`reflect` package) to inspect the Go type of every argument passed into `a...` and match it to your formatting string verb at runtime. This causes interface heap allocations and slows execution compared to `fmt.Println` or direct string concatenations (`+`). For high-throughput loops or performance-critical paths, use explicit conversions (e.g. `strconv.Itoa`) instead of formatting blocks.

## Gotchas

- **No Automatic Newline:** Unlike `Println`, `Printf` **never** appends a newline automatically. If you forget to end your format string with `\n`, your next terminal prints will merge on the same line.
- **Verb/Type Mismatches:** If you supply a type that does not match your formatting verb (e.g. passing a `string` to `%d` or a `struct` to `%s`), Go will not crash. Instead, it prints a raw compilation type warning directly into your output stream:
  ```go
  fmt.Printf("Age: %d\n", "thirty")
  // Outputs: Age: %!d(string=thirty)
  ```
- **Unescaped literal percents:** If you want to print a literal `%` symbol inside a format string, you must write **`%%`**. A single `%` tells the compiler to expect a formatting verb, leading to compile warnings or raw parse errors in the output.

## Related

- struct.basics
- ptr.basics
- os.getenv
