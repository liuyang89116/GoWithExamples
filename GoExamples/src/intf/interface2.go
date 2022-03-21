package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId   int
	basePay int
	pf      int
}

type Contract struct {
	empId    int
	basicPay int
}

func (p Permanent) CalculateSalary() int {
	return p.basePay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicPay
}

func totalExpense(s []SalaryCalculator) {
	sum := 0
	for _, v := range s {
		sum += v.CalculateSalary()
	}
	fmt.Printf("Total expense is $%d\n", sum)
}

func main() {
	emp1 := Permanent{
		empId:   1,
		basePay: 5000,
		pf:      1000,
	}
	emp2 := Permanent{
		empId:   2,
		basePay: 6000,
		pf:      1000,
	}
	emp3 := Contract{
		empId:    3,
		basicPay: 3000,
	}
	employees := []SalaryCalculator{emp1, emp2, emp3}
	totalExpense(employees)
}
