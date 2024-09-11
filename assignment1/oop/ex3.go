package main

import (
	"fmt"
	"math"
)

type Shape interface{
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct{
	Width, Height float64
}

func (r Rectangle)  Area() float64{
	return r.Height* r.Width
}

func PrintArea(s Shape){
	fmt.Printf("The area is %.2f\n",s.Area())
}

func Calculator() {
    
    circle := Circle{Radius: 5}
    rectangle := Rectangle{Width: 4, Height: 6}

    
    PrintArea(circle)      // Prints the area of the circle
    PrintArea(rectangle)   // Prints the area of the rectangle
}