package main

import (
	"fmt"
)

func execpan(){
	defer fmt.Println("I am recovering")
	//panic("This is a panic Suituation due to carona")
	fmt.Println("expecpan completed")
}

func main () {
	execpan()
	fmt.Println("Main block")

}