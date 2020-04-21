// It is also possible to compare a type to an interface.
// If we have a type and if that type implements an interface,
// it is possible to compare this type with the interface it implements.
package main

import "fmt"

type Describer interface {
	Describe()
}
type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func main() {
	findType("Naveen")
	p := Person{
		name: "Naveen R",
		age:  25,
	}
	findType(p)
}

// In the program above, the Person struct implements the Describer interface.
// In the case statement in line no. 22, v is compared to the Describer interface type. p implements Describer and hence this case is satisfied and Describe() method is called.

// This program prints

// unknown type
// Naveen R is 25 years old
