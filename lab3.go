package main

import (
	"fmt"
)

func myAnd (a, b int ){
	c := a & b
	fmt.Printf("\n Bitwise AND: %d & %d = %d", a,b,c)
}

func myOR (a, b int ){
	c := a & b
	fmt.Printf("\n Bitwise OR: %d | %d = %d", a,b,c)
}

func myXOR (a, b int ){
	c := a & b
	fmt.Printf("\n Bitwise AND: %d ^ %d = %d", a,b,c)
}

func myNOT (a, b int ){
	c := a & b
	fmt.Printf("\n Bitwise OR: %d = %d", a,c)
}

//x := 10
//result := x << 3 // 10 * (2^3) = 10 * 8 = 80

func myShift (a, b int){
	shift = b 
	if shift <  0 {

	}

	
	
	
	leftShift = int(a << b )
	rightShift = int(a >> b)
	
	fmt.Printf("\nLeft Shift: %d << %d = %d", a,b, leftShift)
	fmt.Printf("\nRight Shift: %d >> %d= %d" , a,b, rightShift)

		
}


func main() {
	var a, b int

	fmt.Print("Enter the first integer (a): ")
	fmt.Print(&a)
	fmt.Print("Enter the second integer (b): ")
	fmt.Scan(&b)


	
}