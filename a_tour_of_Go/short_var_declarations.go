// Short variable declarations

// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
// Also called a 'walrus operator'
// Outside a function, every statement begins with a keyword (var, func, etc.) and so the := construct is not available.

package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "No!" // notice the walrus operator allows us to omit 'var' and the types are assigned based on the initializer's value

	fmt.Println(i, j, k, c, python, java)
}
