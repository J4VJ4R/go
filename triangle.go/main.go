package main

import (
	"fmt"
	"math"
	// "strconv"
	// "rsc.io/quote"
)



func main(){
	var height, base, area float64

	fmt.Printf("Type height: ")
	fmt.Scanln(&height)
	fmt.Printf("Type base: ")
	fmt.Scanln(&base)

	area = (base * height) / 2
	hyp := math.Hypot(float64(height), float64(base))
	perimeter := float64(height) + float64(base) + hyp
	fmt.Printf("The area is: %.2f \n", area)
	fmt.Printf("The hypotenuse is: %.2f \n", hyp)
	fmt.Printf("The perimeter is: %.2f \n", perimeter)
}