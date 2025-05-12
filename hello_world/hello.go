package main

import (
	"fmt"
)

// defining consts
const (
	french             = "French"
	spanish            = "Spanish"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// hello is a function that returns a string that can be used anywhere
// in the file
// Hello is uppercase for us to use it in different files
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("world", ""))
}

// it is always good to seperate the domain from side effect
// our domain here is the string hello, world and the side effect
// is Println function

// greetingPrefix Function
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
