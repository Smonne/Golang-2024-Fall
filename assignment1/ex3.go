package main

import "fmt"

func control_structures(){
	var num int
	fmt.Printf("Enter an integer: ")
	fmt.Scan(&num)

	if num> 0 {
		fmt.Printf("Positive")
	} else if num< 0 {
		fmt.Printf("Negative" )
	} else {
		fmt.Printf("Zero")
	}

	sum := 0
	for i:=1; i<10; i++ {
		sum+=i
	}
	fmt.Printf("Sum of the first 10 natural numbers : %d\n", sum)

	var day_num int
	fmt.Printf("Enter a number of the day: ")
	fmt.Scan(&day_num)
	switch day:= day_num; day {
	case 1: fmt.Printf("Monday")
	case 2: fmt.Printf("Tuesday")
	case 3: fmt.Printf("Wednesday")
	case 4: fmt.Printf("Thursday")
	case 5: fmt.Printf("Friday")
	case 6: fmt.Printf("Saturday")
	case 7: fmt.Printf("Sunday")
	default: fmt.Println("Invalid input! Please enter a number between 1 and 7.")

	}

	
}
