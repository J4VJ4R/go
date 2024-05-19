package main

import (
	"fmt"
	"rsc.io/quote"
) 	

func main(){
	fmt.Println("Hello world")
	fmt.Println(quote.Go())

	// var firstName, lastName string
	// var age int

	firstName, lastName, age  := "Alex", "Roel", 27
	age = 30
	


	fmt.Println(firstName, lastName, age)
}