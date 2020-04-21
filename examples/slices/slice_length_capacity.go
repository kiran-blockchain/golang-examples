package main

import "fmt"

func main() {
	slice1 := []int{2, 4, 6, 8, 10, 12, 14}
	printSlice(slice1)
	// Slice the slice to give it zero length.
	slice2 := slice1[:0]
	printSlice(slice2)
	// Extend its length.
	slice3 := slice1[:4]
	printSlice(slice3)
	// Drop its first two values.
	slice4 := slice1[2:]
	printSlice(slice4)
}
func printSlice(s []int) {
	fmt.Printf("length =%d capacity=%d %v\n", len(s), cap(s), s)
}
