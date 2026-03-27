// Variables with initializers

// A var declaration can include initializers, one per variable
// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "No!" // Here, Go dynamically assigns TYPES to the variables based on the values we assigned
	fmt.Println(i, j, c, python, java)       // i and j we assigned 'int'. It auto recognizes c and python as BOOLEANS and java as a STRING
}
