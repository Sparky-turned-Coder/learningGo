package main

import "fmt"

func Hello() string { // We created a seperate function, added the keyword 'string'. This means this function returns a 'string'.
	return "Hello, World!"
}

func main() {
	fmt.Println(Hello()) // Here we call our new function, Hello()
}
