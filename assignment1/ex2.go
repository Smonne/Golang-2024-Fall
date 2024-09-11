package main

import "fmt"

func declaration(){
	var year int = 20
	var height float64 = 6.6
	var name = "Timothee Shalame"
	var isStudent bool = false

	
	country:= "USA"
	isActor := true

	fmt.Printf("year: %d, type: %T\n", year, year)
	fmt.Printf("height: %.1f, type: %T\n", height, height)
	fmt.Printf("name: %s, type: %T\n", name, name )
	fmt.Printf("isStudent: %t, type: %T\n", isStudent, isStudent)

	fmt.Printf("country: %s, type: %T\n", country, country)
	fmt.Printf("isActor: %t, type: %T\n", isActor, isActor)
}