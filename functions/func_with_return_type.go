package main

import "fmt"

// Function with int as return type
func add(x int, y int) int {
	total := 0
	total = x + y
	return total
}

//named returns
func math(x int, y int) (subtracted int, mulitply int, divide int) {
	subtracted = x - y
	mulitply = x * y
	divide = x / y

	return
}

func main() {
	// Accepting return value in varaible
	sum := add(20, 30)
	sub, mul, div := math(80, 50)
	fmt.Println(sum)
	fmt.Println(sub)
	fmt.Println(mul)
	fmt.Println(div)
}
