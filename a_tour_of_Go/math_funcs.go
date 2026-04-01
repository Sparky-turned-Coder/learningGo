package main

import "fmt"

//func main() {
//	sum := 0
//	for i := 0; i < 10; i++ {
//		sum += i
//		fmt.Println(i)
//	}
//	fmt.Println(sum)
//}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func divide(x, y int) int {
	return x / y
}

func modulo(x, y int) int {
	return x % y
}

func multiply(x, y int) int {
	return x * y
}

func main() {
	fmt.Printf("The sum of 6 and 7 is %v\n", add(6, 7))
	fmt.Printf("The difference of 6 and 7 is %v\n", subtract(6, 7))
	fmt.Printf("The quotient of 6 and 7 is %v\n", divide(6, 7))
	fmt.Printf("The remainder of 6 by 7 is %v\n", modulo(6, 7))
	fmt.Printf("The product of 6 and 7 is %v\n", multiply(6, 7))
}
