package main

import "GoExamples/src/oop/employee"

func main() {
	e := employee.New("L", "Messi", 30, 5)
	e.LeavesRemaining()
}
