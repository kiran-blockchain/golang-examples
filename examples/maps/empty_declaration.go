package main

import "fmt"

func main() {
	var employee = map[string]int{}
	fmt.Println(employee)        // map[]
	fmt.Printf("%T\n", employee) // map[string]int
}
