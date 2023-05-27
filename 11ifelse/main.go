package main

import "fmt"

func main() {
	age := 16

	if age < 10 {
		fmt.Println("Age is less than 10")
	} else if age >= 10 && age < 18 {
		fmt.Println("Age is between 10 and 18")
	} else {
		fmt.Println("You are an adult")
	}

	// One other way to using if and else
	if age := 10; age < 18 {
		fmt.Println("You are not an adult")
	} else {
		fmt.Println("You are an adult")
	}
}
