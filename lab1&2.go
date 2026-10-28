package main

import "fmt"

func main() {
    var a , b int 
    
    fmt.Print("Enter (a) value: ")
    fmt.Scan(&a)
    fmt.Print("Enter (b) value: ")
    fmt.Scan(&b)

    fmt.Printf("\nValue (a)= %d, Value (b)= %d", a,b)

    a += b
    fmt.Printf("\na + b = %d", a )

    a -= b
    fmt.Printf("\na - b = %d", a )

    a *= b
    fmt.Printf("\na * b = %d", a )

    a /= b
    fmt.Printf("\na / b = %d", a )

    a %= b
    fmt.Printf("\na % b = %d\n", a)



    // Logical operator

    positive := (a > 0) && (b > 0)
	fmt.Printf("\nBoth positive (%d > 0 && %d > 0): %t", a, b, positive)

	
	greater := (a > b) || (b > a)
	fmt.Printf("\nOne greater than the other (%d > %d || %d > %d): %t", a, b, b, a, greater)

	
	notEqual := !(a == b)
	fmt.Printf("\nNot equal !(%d == %d): %t\n", a, b, notEqual)
  
}

