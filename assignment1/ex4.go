package main

import "fmt"


func add(a int, b int) int {
    return a + b
}


func swap(str1, str2 string) (string, string) {
    return str2, str1
}


func quotientAndRemainder(dividend, divisor int) (int, int) {
    quotient := dividend / divisor
    remainder := dividend % divisor
    return quotient, remainder
}

func functions() {
   
    sum := add(5, 3)
    fmt.Printf("Sum of 5 and 3: %d\n", sum)

    
    first, second := swap("Hello", "World")
    fmt.Printf("Swapped strings: %s, %s\n", first, second)

    
    quotient, remainder := quotientAndRemainder(10, 3)
    fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)
}