# cmd.test

## Card

Run unit tests and benchmarks in your Go project.

```bash
# Run all tests in the current package
go test .

# Run all tests in the project recursively (highly recommended!)
go test -v ./...
```

### Subcard: Run Specific Test

```bash
go test -run TestParseHeading -v .
```

### Subcard: Check Test Coverage

```bash
go test -cover .
```

## Example

Walkthrough of writing, running, and analyzing coverage for a standard unit test in Go:

1. Create your source file `math.go`:
```go
package main

func Double(x int) int {
	return x * 2
}
```

2. Create your unit test file `math_test.go` ending in `_test.go`:
```go
package main

import "testing"

func TestDouble(t *testing.T) {
	got := Double(3)
	want := 6
	if got != want {
		t.Errorf("Double(3) = %d; want %d", got, want)
	}
}
```

3. Run tests in verbose mode to see step-by-step executions:
```bash
go test -v .

# Expected Output:
# === RUN   TestDouble
# --- PASS: TestDouble (0.00s)
# PASS
# ok      myapp   0.114s
```

4. Run the test suite and output coverage statistics:
```bash
go test -cover .

# Expected Output:
# ok      myapp   0.125s   coverage: 100.0% of statements
```

## Deep

Go has a first-class built-in testing framework. Test files must end in `_test.go` and live in the same directory as the package they are testing. Test functions must start with `Test` and have the signature `func TestX(t *testing.T)`.

## Gotchas

- **Cached Test Results:** Go caches successful test results. If you change nothing, Go will output `(cached)` and won't re-run tests. To force-run tests and bypass cache, use the `-count=1` flag:
  ```bash
  go test -count=1 ./...
  ```

## Related

- cmd.build
- setup.project
