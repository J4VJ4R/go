package main

import (
	"fmt"
	// "math"
	"strconv"
	// "rsc.io/quote"
)



func main(){
	// Variable definition
	// firstName, lastName, age := "Javier", "Jaramillo", 41
	// age = 42

	const Pi float32 = 3.14

	const (
		x = 100
		y = 0b1010
		z = 0o22
		w = 0xff
	)
	
	const (
		Monday = iota + 1
		Tuesday
		Wednsday
		Thursday
	)
	
	
	fmt.Println(Pi)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(w)
	fmt.Println(Tuesday)

	s := "100"
	i, _ := strconv.Atoi(s)
	fmt.Println(i + i)

	y := 42
	a := strconv.Itoa(y)
	fmt.Println(a)
	
	var name string
	var age int

	println("Type your name: ")
	fmt.Scanln(&name)
	println("Type your age: ")
	fmt.Scanln(&age)
	fmt.Printf("My name is %s and my age is %d .", name, age)

}