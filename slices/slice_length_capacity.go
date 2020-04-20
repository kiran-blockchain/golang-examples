package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 14} //array
	slice1 := arr[0:7]
	fmt.Println("-----")
	fmt.Printf("%p", &arr)
	fmt.Println("")
	fmt.Printf("%p", &slice1)
	fmt.Println("")
	slice1[3] = 30000
	fmt.Println("")
	slice1 = append(slice1, 1)
	fmt.Println("--after 1---")
	fmt.Printf("%p", &arr)
	fmt.Println("")
	fmt.Printf("%p", &slice1)
	fmt.Println("")
	slice1 = append(slice1, 1900)
	fmt.Println("--after 1900---")
	fmt.Printf("%p", &arr)
	fmt.Println("")
	fmt.Printf("%p", &slice1)

	slice2 := append(slice1, 2000)
	fmt.Println("--after 2000---")
	fmt.Printf("%p", &arr)
	fmt.Println("")
	fmt.Printf("%p", &slice2)
	//	s[0] = 1000
	//output
	fmt.Println(arr)
	fmt.Println(slice1)
	//fmt.Println(s)
}
func printSlice(s []int) {
	fmt.Printf("length =%d capacity=%d %v\n", len(s), cap(s), s)
}
