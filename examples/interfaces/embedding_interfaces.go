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

type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
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
	var empOp EmployeeOperations = e
	empOp.DisplaySalary()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())
}

// Any type is said to implement EmployeeOperations interface
// if it provides method definitions for the methods present in
// both SalaryCalculator and LeaveCalculator interfaces.

// The Employee struct implements EmployeeOperations interface
// since it provides definition for both DisplaySalary and
// CalculateLeavesLeft methods in lines 29 and 33 respectively.

// In line 46, e of type Employee is assigned to empOp of type EmployeeOperations.
// In the next two lines, the DisplaySalary() and CalculateLeavesLeft() methods are called on empOp.

// This program will output

// Naveen Ramanathan has salary $5200
// Leaves left = 25
