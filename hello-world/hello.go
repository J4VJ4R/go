package main

import (
	"fmt"
	"rsc.io/quote"
) 	
const Pi  = 3.14
const (
	x = 100
	y = 0b1010
	z = 0o12
	w = 0xFF
)
const(
	Sunday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	
)
func main(){
	fmt.Println("Hello world")
	fmt.Println(quote.Go())

	// var firstName, lastName string
	// var age int

	firstName, lastName, age  := "Alex", "Roel", 27
	age = 30
	


	fmt.Println(firstName, lastName, age)
	fmt.Println(Pi)
	fmt.Println(x, y, z, w)
	fmt.Println(Sunday)
}