package main

import "fmt"

// Define a constant
const englishHelloPrefix = "Hello, "

func Hello(name string) string { // We created a seperate function, added the keyword 'string'. This means this function returns a 'string'.
	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("world")) // Here we call our new function, Hello()
}
