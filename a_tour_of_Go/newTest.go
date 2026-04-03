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
}
