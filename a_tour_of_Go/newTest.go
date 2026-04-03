package main

import (
	"fmt"
	"strconv"
)

func concat(name1, name2 string) string {
	return name1 + " " + name2
}

func addition(x, y int) int {
	return x + y
}

func convert(myString string) int {
	x, err := strconv.Atoi(myString)
	if err != nil || x < 1 {
		fmt.Println("Conversion failed.")
	}
	return x
}

const X = 1 << 6

func main() {
	word := "Pleasant"
	wordy := "Day"
	num1 := 31
	num2 := 21
	x := "17"
	y := "18"
	z := "155"
	fmt.Printf("Well, hello there! %v isn't it?\n", concat(word, wordy))
	fmt.Printf("The sum of your two numbers is: %v\n", addition(num1, num2))
	fmt.Printf("The numerical representation of the strings '%v', '%v', and '%v' is %v, %v, %v, respectively.\n", x, y, z, convert(x), convert(y), convert(z))
	fmt.Printf("The complex conversion of '1 << 6' is: %v\n", complex128(X))
	fmt.Printf("The integer conversion of '1 <<6' is: %v\n", int(X))
	fmt.Printf("The integer conversion plus 1 of '1 << 6' is: %v\n", int(X+1))
	fmt.Printf("The float conversion plus .55 of '1 << 6' is: %v\n", float64(X+.55))
}
