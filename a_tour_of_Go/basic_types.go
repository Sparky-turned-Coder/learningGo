// Go's Basic Types

/*

*** See notes.md for more details on basic types

Go's basic types are:

- bool

- string

- int int8 int16 int32 int64  (default to using 'int' unless a specific use-case requires direct sizing of a signed integer)
- uint uint8 uint16 uint32 uint64 uintptr (unsigned)

- byte > alias for uint8  (8 bits == byte)
- rune > alias for int32 (a 32-bit signed integer and therefore occupies 4 bytes in memory)
	   > represents a Unicode code point

- float32 float64

- complex64 complex128
*/

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T | Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T | Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T | Value: %v\n", z, z)
}
