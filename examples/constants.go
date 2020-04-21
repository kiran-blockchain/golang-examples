//A constant const contains data which is not changed
// Syntax:- const identifier [type] = value

package main

import "fmt"

func main() {
	const HEIGHT int = 100
	const WIDTH int = 200
	var area int
	area = HEIGHT * WIDTH
	fmt.Printf("value of area : %d", area)
}
