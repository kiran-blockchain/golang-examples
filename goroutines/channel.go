package main

import "fmt"

func sum(a []int, c chan<- string) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Println(sum)
	c <- "Hey I am sum 1" // ? Assinging the sum to Channel C
	fmt.Println("length of the channel is: ", len(c))
	fmt.Println("cap of the channel is: ", cap(c))
}
func sum2(a []int, c chan string) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Println(sum)
	
	fmt.Println("length of the channel is: ", len(c))
	fmt.Println("cap of the channel is: ", cap(c))
}

func main() {
	a := []int{1, 2, 2}
	b := []int{1, 56, 78}
	c := make(chan string, 2)

	go sum(a, c)
	go sum2(b, c)

	fmt.Println("length of the channel is: ", len(c))
	fmt.Println("cap of the channel is: ", cap(c))
	x,err:= <-c
	y,err2 :=<-c
	
	if err = nil {
	fmt.Println("In main ", x)
	}
	//fmt.Println("In main ", y)
	// fmt.Println("In main ", z)
	// fmt.Println("In main ", hh)
}
