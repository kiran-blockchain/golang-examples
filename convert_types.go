package main

import (
	"fmt"
	"strconv"
)

func main() {

	//String to Int
	strVar := "100"
	intVar, _ := strconv.Atoi(strVar)

	strVar1 := "-52541"
	intVar1, _ := strconv.ParseInt(strVar1, 10, 32)

	strVar2 := "101010101010101010"
	intVar2, _ := strconv.ParseInt(strVar2, 10, 64)

	//String to Float
	s := "3.1415926535"
	f, _ := strconv.ParseFloat(s, 8)
	fmt.Printf("%T, %v\n", f, f)

	s1 := "-3.141"
	f1, _ := strconv.ParseFloat(s1, 8)
	fmt.Printf("%T, %v\n", f1, f1)

	s2 := "-3.141"
	f2, _ := strconv.ParseFloat(s2, 32)
	fmt.Printf("%T, %v\n", f2, f2)

	//String to Boolean
	str1 := "true"
	b1, _ := strconv.ParseBool(str1)
	fmt.Printf("%T, %v\n", b1, b1)

	str2 := "t"
	b2, _ := strconv.ParseBool(str2)
	fmt.Printf("%T, %v\n", b2, b2)

	str3 := "0"
	b3, _ := strconv.ParseBool(str3)
	fmt.Printf("%T, %v\n", b3, b3)

	str4 := "F"
	b4, _ := strconv.ParseBool(str4)
	fmt.Printf("%T, %v\n", b4, b4)

}
