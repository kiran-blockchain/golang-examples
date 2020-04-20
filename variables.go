package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	var c complex128
	fmt.Printf("%T %T %T %T   %T\n", i, f, b, s, c)          // Prints type of the variable
	fmt.Printf("%v   %v      %v  %q    %q\n", i, f, b, s, c) //prints initial value of the variable
}
