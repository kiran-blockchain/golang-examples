package main

import "fmt"

type Address interface  {
	getName(name string)
	getAge(age  int)
}
type Person struct {
	address Address
}

func main(){
	a := Person{}
	a.getAge()
}

