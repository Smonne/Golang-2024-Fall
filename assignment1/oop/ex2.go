package main

import "fmt"

type Employee struct{
	Name string
	Id int
}

type Manager struct{
	Employee
	Department string
}

func (e Employee) Work(){
	fmt.Printf("The employee's name %s. ID is %d.\n", e.Name, e.Id)
}

func Office(){
	manager:= Manager{
		Employee: Employee{Name: "Timothee SHalamet", Id: 2314},
		Department: "Marketing",
	}
	manager.Work()

	fmt.Printf("The manager's name %s works in the Department of %s.\n", manager.Name, manager.Department)
}