package main

import (
	"fmt"
	"math/rand"
)

func main(){
	option()
}
func game(){
	randNumber := rand.Intn(100)
	var turns int
	var typeNumber int
	const maxTurns = 10
	
	for turns < maxTurns {
			fmt.Printf("================================= \n")
			fmt.Printf("You have %d turns \n", maxTurns - turns)
			fmt.Println("Type your number: ")
			fmt.Scanln(&typeNumber)

			if typeNumber < randNumber{
				fmt.Printf("The number is less \n")
				turns++
			}else if typeNumber > randNumber{
				fmt.Printf("The number is greater \n")
				turns++
			}else if typeNumber == randNumber{
				fmt.Printf("Congratulations you win!! \n")
				option()
				return
			}
			
	}
}
func option(){
	var option string 
	fmt.Printf("=========THE GAME NUMBER========= \n")
	fmt.Println("Play the Game Number s/n")
	fmt.Scanln(&option)

	switch option {
	case "s":
		game()
	case "n":
		fmt.Println("Thanks see you soon!!")
		break
	}
	
}