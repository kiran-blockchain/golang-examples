package main

import (
	"fmt"
)

type Worker interface {
	Work() // Worker interface has one method Work()
}

type Person struct {
	name string
}

// Person struct type implements that interface
func (p Person) Work() {
	fmt.Println(p.name, "is working")
}

//The describe function in line no.17 prints the value and concrete type of the interface.
func describe(w Worker) {
	fmt.Printf("Interface type %T value %v\n", w, w)
}

func main() {
	p := Person{
		name: "Naveen",
	}
	var w Worker = p //we assign the variable p of type Person to w which is of type Worker
	//The concrete type of w is Person and it contains a Person with name field Naveen.
	describe(w)
	w.Work()
}

// Output:-
// Interface type main.Person value {Naveen}
// Naveen is working
