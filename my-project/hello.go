package main

import "fmt"

func Hello(name string) string { // We created a seperate function, added the keyword 'string'. This means this function returns a 'string'.
	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(Hello("world")) // Here we call our new function, Hello()
}
