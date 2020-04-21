package main

import "fmt"

func main() {

	// Defer executes a function after the inclosing function finishes
	// Defer can be used to keep functions together in a logical way
	// but at the same time execute one last as a clean up operation
	// Ex. Defer closing a file after we open it and perform operations

	defer printTwo()
	printOne()

	// Use recover() to catch a division by 0 error

	fmt.Println(Div(3, 0))
	fmt.Println(Div(3, 2))

	// We can catch our own errors and recover with panic & recover

	demPanic()

}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

// Used to demonstrate defer

func printOne() { fmt.Println(1) }

func printTwo() { fmt.Println(2) }

// If an error occurs we can catch the error with recover and allow
// code to continue to execute

func Div(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	solution := num1 / num2
	return solution
}

// Demonstrate how to call panic and handle it with recover

func demPanic() {

	defer func() {

		// If I didn't print the message nothing would show

		fmt.Println(recover())

	}()
	panic("PANIC")

}
