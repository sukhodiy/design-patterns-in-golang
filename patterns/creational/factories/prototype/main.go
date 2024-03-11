package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

type EmployeeRole int

const (
	Developer EmployeeRole = iota
	Manager
)

func NewEmployee(role EmployeeRole) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "developer", 60000}
	case Manager:
		return &Employee{"", "manager", 80000}
	default:
		panic("unsupported role")
	}
}

func main() {
	m := NewEmployee(Manager)
	m.Name = "Sam"
	fmt.Println(m)
}
