package main

import "fmt"

func main() {
	var username string = "Varun"
	fmt.Println(username)
	fmt.Printf("Varun is of type %T \n",username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Varun is of type %T \n",isLoggedIn)

	var number int = 2345
	fmt.Println(number)
	fmt.Printf("Number is of type %T \n",number)

	var floatNumber float64 = 32325.3232333
	fmt.Println(floatNumber)
	fmt.Printf("floatNumber is of type %T \n",floatNumber)

	// implicite type
	var sentence = "How are you?"
	fmt.Println(sentence)

	// no var style
	users := 2000000
	fmt.Println(users)
	fmt.Printf("The type of users is %T \n",users)

	// Constant
	const PI = 3.14
	fmt.Println(PI)
	fmt.Printf("Type of PI is %T \n",PI)
}