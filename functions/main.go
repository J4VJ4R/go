package main

import "fmt"

func main(){
	// name := "Javi2"
	fmt.Println(helloworld("Javi3"))
	fmt.Println(cal(10, 2))
}

func helloworld(name string)string{
	return "Hola " + name
}
func cal(num1, num2 int)(sum, mul int){
	sum = num1 + num2
	mul = num1 * num2
	return sum, mul
}