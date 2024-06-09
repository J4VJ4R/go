package main

import "fmt"

func main(){
	number := 5
	if number > 0{
		fmt.Println("The number is positive.")
	}else{
		fmt.Println("The number is negative.")

	}

	number = 1

	switch number {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
		
	}

}