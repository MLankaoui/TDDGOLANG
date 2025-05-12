# Writing tests
1. name the file ```xxx_test.go```
2. the function must start with ```Test```
3. test function takes one argument ```t *testing.T```
4. to use ``` *testing.T``` type, we need to ```import "testing```

# defining consts:
```go
const englishHelloPrefix = "Hello, "
```

## Subsets:
- We add subsets to describe different scenarios.

- **Example**
```go
    t.Run("saying hello to people", func(t *testing.T) {
		// what you got from the function
		got := Hello("Chris")
		// what we expect
		want := "Hello, Chris"

		// checking if what we got is what we expected
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, World when an empty string is supplied", func(t *testing.T) {
		got := Hello("")

		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
```