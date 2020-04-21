package main

import (
	"fmt"
)

func main() {
	// number := 10
	// squareNum := func() int {
	// 	number *= number
	// 	return number
	// }
	// fmt.Println(squareNum())
	// fmt.Println(squareNum())
	x := fibSeries()
	for i:=0; i<5;i++ {
		fmt.Println(x())
	}
}

func fibSeries () func()int{
	f1:=0
	f2:=1
	return func ()int{
		f2,f1 = (f1+f2),f2
		return f1
	}
}
