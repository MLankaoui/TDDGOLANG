package main

// importing the testing package
import "testing"

func TestHello(t *testing.T) {
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

}
