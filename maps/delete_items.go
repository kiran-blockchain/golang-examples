package main

import "fmt"

func main() {
	var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20
	employee["Rocky"] = 30
	employee["Josef"] = 40

	fmt.Println(employee)

	delete(employee, "Mark")
	fmt.Println(employee)
}
