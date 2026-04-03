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

func convert(num1 string) int {
	new_num, err := strconv.Atoi(num1)

	if err != nil {
		fmt.Println("Conversion failed.")
	}
	return new_num
}

func main() {
	word := "Pleasant"
	wordy := "Day"
	x := 17
	y := 18
	z := "1"
	fmt.Printf("Well, hello there! %v isn't it?\n", concat(word, wordy))
	fmt.Printf("The sum of your two numbers is: %v\n", addition(x, y))
	fmt.Printf("The numerical representation of string '1' is: %v\n", convert(z))
}
