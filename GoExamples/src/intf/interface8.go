package main

import "fmt"

type SalaryCalculator2 interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateRemainingLeaves() int
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
	fmt.Printf("%s %s has salary $%d\n", e.firstName, e.lastName, e.basicPay+e.pf)
}

func (e Employee) CalculateRemainingLeaves() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	e := Employee{
		firstName:   "L",
		lastName:    "Messi",
		basicPay:    1000,
		pf:          1500,
		totalLeaves: 20,
		leavesTaken: 5,
	}

	var s SalaryCalculator2 = e
	s.DisplaySalary()
	var l LeaveCalculator = e
	fmt.Println("leaves left =", l.CalculateRemainingLeaves())
}
