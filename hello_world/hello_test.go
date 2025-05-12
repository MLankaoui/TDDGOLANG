package main

// importing the testing package
import "testing"

func TestHello(t *testing.T) {
	// what you got from the function
	got := Hello()
	// what we expect
	want := "Hello, worl"

	// checking if what we got is what we expected
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
