package main

import "fmt"

type Person struct{
	Name string
	Age int
}

func (p Person) Greet(){
	fmt.Printf("Hello, my name is %s and I'm %d years old\n", p.Name, p.Age)
}

func main() {
	person := Person{Name: "Timothee Shalamet", Age: 28}
	person.Greet()
}