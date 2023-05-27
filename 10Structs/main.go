package main

import "fmt"

func main() {

	// how to define the structure in golang
	type User struct {
		Name     string
		Age      int
		Email    string
		Verified bool
	}

	user1 := User{"Varun", 25, "varun@gmail.com", true}

	fmt.Println(user1)
	fmt.Printf("The details of the user1 are: %+v \n", user1)

	fmt.Printf("Name of the user: %v\n", user1.Name)
	fmt.Printf("Email of the user: %v\n", user1.Email)

}
