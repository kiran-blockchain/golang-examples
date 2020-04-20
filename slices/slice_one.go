package main

import "fmt"

func main() {
	arr := []string{"I", "am", "an", "array"}

	fmt.Println(arr)

	mySlice := arr[0:3]
	mySlice[1] = "kiran"

	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println(arr)
}
