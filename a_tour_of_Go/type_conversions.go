/* Type conversions

The expression T(v) converts the value(v) to the type(T)

Some numeric conversions:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

we converted i to a float and stored it inside f.
we converted f to a uint and stored it inside u.

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

/*

What happens:

- we assign 3 to 'x' and 4 to 'y'
- then we take the square root of (3*3 + 4*4) and convert it to a float stored in 'f'(5.0)
- then we convert f(5.0) to an unsigned integer stored in 'z'

Therefore, the output is:  3, 4, 5

*/
