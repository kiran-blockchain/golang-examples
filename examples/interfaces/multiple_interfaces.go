package main

import (
	"fmt"
)

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	e := Employee{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var s SalaryCalculator = e
	s.DisplaySalary()
	var l LeaveCalculator = e
	fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())
}

// The program above has two interfaces SalaryCalculator and LeaveCalculator
// declared in lines 7 and 11 respectively.

// The Employee struct defined in line no. 15 provides
// implementations for the DisplaySalary method of SalaryCalculator interface
// in line no. 24 and the CalculateLeavesLeft method of LeaveCalculator interface in line no. 28.
// Now Employee implements both SalaryCalculator and LeaveCalculator interfaces.

// In line no. 41 we assign e to a variable of type SalaryCalculator interface
// and in line no. 43 we assign the same variable e to a variable of
// type LeaveCalculator. This is possible since e which of type Employee
// implements both SalaryCalculator and LeaveCalculator interfaces.

// This program outputs,

// Naveen Ramanathan has salary $5200
// Leaves left = 25
