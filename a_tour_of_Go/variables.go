// Variables

// The 'var' statement declares a list of variables; as in function argument lists, the type is last.
// A 'var' statement can be at package or function level. We see both in this example.

package main

import "fmt"

var c, python, java bool // var statement assigned at package level

func main() {
	var i int // var statement assigned at function level
	fmt.Println(i, c, python, java)
}

// Output is:   0 false false false
// Remember, this is because all variables by default have a value of 0 until you assign something else to them.
// In the case of a boolean, 0 translates to FALSE
