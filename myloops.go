package main

import (
	"fmt"
)

func main() {
	x := []int{1, 8, 9}
	sum := 0
	for _, val := range x {
		sum = sum + val
		fmt.Println(sum)
	}
}
