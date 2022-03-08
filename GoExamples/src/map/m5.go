package main

import "fmt"

type employee struct {
	age     int
	country string
}

func main() {
	emp1 := employee{
		age:     26,
		country: "CN",
	}
	emp2 := employee{
		age:     30,
		country: "USA",
	}
	emp3 := employee{
		age:     26,
		country: "FRA",
	}

	employeeMap := map[string]employee{
		"Bob": emp1,
		"Amy": emp2,
		"Jay": emp3,
	}

	for name, info := range employeeMap {
		fmt.Printf("Employee: name[%s], age[%d], country[%s]\n",
			name, info.age, info.country)
	}
}
