package main

import ("fmt"
		"math"				// THE WAY OF IMPORTING HEADER FILES OR LIBRARY IN GO
		"time")		


func main() {

	fmt.Println("Welcome! The Go Programming Language is tuts")

	fmt.Println("The Square root of 25 and 36 is", math.Sqrt(25), math.Sqrt(36))
	
	fmt.Println("What is the time right now ?  ", time.Now())

}