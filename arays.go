package main

import (
	"fmt"
)

func main() {
	x := []int{1, 99, 857, 765, 987}
	//	x[3] = "1000"
	//x[5] = 900
	fmt.Println(x)
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

}
