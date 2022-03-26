package main

import "fmt"

type Describer interface {
	Describe()
}

type Person2 struct {
	name string
	age  int
}

func (p Person2) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() {
	fmt.Printf("state %s, country %s\n", a.state, a.country)
}

func main() {
	var d1 Describer
	p1 := Person2{
		name: "Messi",
		age:  25,
	}
	d1 = p1
	d1.Describe()

	p2 := Person2{
		name: "Bob",
		age:  18,
	}
	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := Address{
		state:   "WA",
		country: "USA",
	}
	d2 = &a
	d2.Describe()
}
