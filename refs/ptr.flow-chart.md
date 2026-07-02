# ptr.flow-chart

## Card

An ASCII decision tree flow chart for choosing between Pointer (`*T`) vs. Value (`T`) in Go.

```txt
                Is the type a map, slice, channel, or function?
                                  /       \
                               Yes         No
                               /             \
                   Use VALUE [T]              Does the method/function
               (They are headers that          need to modify fields?
                reference data anyway)                /       \
                                                    Yes         No
                                                    /             \
                                        Use POINTER [*T]           Is it a large struct
                                                                  or has a sync.Mutex?
                                                                        /       \
                                                                     Yes         No
                                                                     /             \
                                                         Use POINTER [*T]           Use VALUE [T]
                                                       (Avoids copy overhead)    (Safe & lightweight)
```

### Subcard: Simple Rules

- **Use Pointers (`*T`)** if you need to mutate, if the struct is large, or if it contains locks/connections.
- **Use Values (`T`)** for basic primitives, small read-only structs, maps, slices, and channels.

## Example

```go
// 1. Modifying state -> Pointer receiver
func (u *User) UpdateEmail(email string) { u.Email = email }

// 2. Heavy struct -> Pointer receiver (efficiency)
type Heavy struct { Data [1024]byte }
func (h *Heavy) Process() {}

// 3. Small read-only -> Value receiver
type Point struct { X, Y int }
func (p Point) Distance() int { return p.X + p.Y }
```

## Deep

Choosing between pointers and values is a critical performance and correctness decision in Go:
- **Receiver Guidelines:** Once you choose pointers for a struct's receiver type, use pointers consistently across all methods on that type, even those that are read-only.
- **Slices and Maps:** Slices, maps, and channels are already internal reference structures (pointing to underlying arrays/hash tables). You do not need `*[]int` or `*map[string]string` because passing them as values already allows modifying their elements.

## Gotchas

- **Mixing Receivers:** Mixing value and pointer receivers on the same struct type can lead to confusion and bugs, as some methods will mutate the original while others silently fail to do so because they work on a copy.

## Related

- ptr.basics
- struct.methods
- ptr.pkg-usage
