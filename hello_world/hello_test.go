package main

// importing the testing package
import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		// what you got from the function
		got := Hello("Chris", "Spanish")
		// what we expect
		want := "Hola, Chris"

		// checking if what we got is what we expected
		assetCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")

		want := "Hello, World"

		assetCorrectMessage(t, got, want)
	})

}

// we have repeated code so let's put it on a seperate function
func assetCorrectMessage(t testing.TB, got, want string) {
	// it helps specify where the test really failed insted of it being inside our helper function
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
