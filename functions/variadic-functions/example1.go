package main

import "fmt"

func main() {
	x("red", "blue", "green", "yellow")
}

func x(s ...string) {
	fmt.Println(s[0])
	fmt.Println(s[3])
}
