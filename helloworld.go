package main

import (
	"fmt"
)

func main() {
	// var myName string
	// myName = "Kiran Paturi"

	// var myName2 float64 = 3.4576
	// fmt.Println(myName)
	// //fmt.Println(reflect.TypeOf(myName2))
	// fmt.Printf("%T \n", myName2)
	//fmt.Println("I am zero")
	defer fmt.Println("I am deferred1")
	defer fmt.Println("I am deferred2")
	defer fmt.Println("I am deferred3")
	//fmt.Println("I am first")
}
