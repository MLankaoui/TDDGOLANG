package main

import "fmt"

// defining consts
const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// hello is a function that returns a string that can be used anywhere
// in the file
// Hello is uppercase for us to use it in different files
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == french {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}

// it is always good to seperate the domain from side effect
// our domain here is the string hello, world and the side effect
// is Println function
