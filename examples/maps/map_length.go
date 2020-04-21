package main

import "fmt"

func main() {
	var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20

	// Empty Map
	employeeList := make(map[string]int)

	fmt.Println(len(employee))     // 2
	fmt.Println(len(employeeList)) // 0
}
