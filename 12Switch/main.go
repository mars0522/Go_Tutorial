package main

import (
	"fmt"
	"math/rand"
)

func main() {
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("You rolled 1")
	case 2:
		fmt.Println("You rolled 2")
	case 3:
		fmt.Println("You rolled 3")
	case 4:
		fmt.Println("You rolled 4")
	case 5:
		fmt.Println("You rolled 5")
	case 6:
		fmt.Println("You rolled 6")
	default:
		fmt.Println("Whats that!")
	}
}
