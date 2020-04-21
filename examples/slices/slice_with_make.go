package main

import "fmt"

func main() {
	slice := make([]int, 10)
	printSlice("slice", slice)
	slice1 := make([]int, 0, 10)
	printSlice("slice1", slice1)
	slice2 := slice1[:5]
	printSlice("slice2", slice2)
	slice3 := slice2[2:5]
	printSlice("slice3", slice3)
}
func printSlice(s string, x []int) {
	fmt.Printf("%s length=%d capacity=%d %v\n",
		s, len(x), cap(x), x)
}
