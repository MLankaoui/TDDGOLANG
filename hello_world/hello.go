package main

import "fmt"

// defining consts
const englishHelloPrefix = "Hello, "

// hello is a function that returns a string that can be used anywhere
// in the file
// Hello is uppercase for us to use it in different files
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}

// it is always good to seperate the domain from side effect
// our domain here is the string hello, world and the side effect
// is Println function
