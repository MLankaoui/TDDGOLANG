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

## In case of repeated checking:
- It is better to refactor the code to avoid repetition:
- **Example**
```go
// we have repeated code so let's put it on a seperate function
func assetCorrectMessage(t testing.TB, got, want string) {
	// it helps specify where the test really failed insted of it being inside our helper function
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
```
## Discipline:
- the cycle
1. Write a test
2. Make the compiler pass
3. Run the test, see that it fails and check the error message is meaningful
4. Write enough code to make the test pass
5. Refactor

## Switch
- when we have lots of ```go if``` statements it is better to use ```go switch```
- **Example**
```go
switch language {
case spanish:
	prefix = spanishHelloPrefix
case french:
	prefix = frenchHelloPrefix
}
```