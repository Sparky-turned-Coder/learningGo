package main

import (
	"fmt"
	"strings"
)

func addition(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("Division by zero attempted.")
	}
	return x / y, nil
}

func numPrompt(prompt string) float64 {
	var num float64

	fmt.Printf("%s", prompt)
	fmt.Scanln(&num)
	return num
}

func main() {
	for {
		x := numPrompt("Enter the first number: ")
		y := numPrompt("Enter the second number: ")
		var op string
		var result float64
		var err error

		fmt.Printf("Choose your operation (+, -, *, /): ")
		fmt.Scanln(&op)

		switch op {
		case "+":
			result = addition(x, y)
		case "-":
			result = subtract(x, y)
		case "*":
			result = multiply(x, y)
		case "/":
			result, err = divide(x, y)
		default:
			fmt.Println("That is not a valid operation.")
			return
		}
		if err != nil {
			fmt.Println("Error: %v", err)
			return
		}
		fmt.Printf("The result of your chosen operation is: %.2f\n", result)

		// Below, we turned this main function into a loop that the user can choose to perform over and over again.
		var again string
		fmt.Printf("Do you want to perform another calculation? (y/n): ")
		fmt.Scanln(&again)

		again = strings.ToLower(again)
		if again != "y" && again != "yes" {
			break
		}
	}
}

/*

Bottom line: You’ve built a fully functional, user-friendly CLI calculator in Go that demonstrates:

 Functions
 Conditionals (switch)
 Error handling (error)
 User input/output
 Type handling (float64)
*/

// If/else statement we replaced with the Switch statement above
//
// if op == "+" {
// 	fmt.Printf("The sum of your two numbers is: %v\n", addition(num1, num2))
// } else if op == "-" {
// 	fmt.Printf("The difference of your two numbers is: %d\n", subtract(num1, num2))
// } else if op == "*" {
// 	fmt.Printf("The product of your two numbers is: %d\n", multiply(num1, num2))
// } else if op == "/" {
// 	quotient, err := divide(num1, num2)
// 	if err != nil {
// 		fmt.Printf("Error occurred: %v\n", err)
// 		return
// 	}
// 	fmt.Printf("The quotient of your two numbers is: %d\n", quotient)
// } else {
// 	fmt.Println("That is not a valid operation.")
// }
