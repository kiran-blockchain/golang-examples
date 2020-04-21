package main

import "fmt"

func main() {
	slice1 := []int{2, 4, 8, 10, 12, 14}
	slice2 := slice1[2:4]
	fmt.Println(slice2)
	slice3 := slice1[:3]
	fmt.Println(slice3)
	slice4 := slice1[2:]
	fmt.Println(slice4)
	fmt.Println(slice1)
}
