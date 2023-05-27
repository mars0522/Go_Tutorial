package main

import "fmt"

func main() {


	// 1st way to create an array
	var fruits [4]string
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[3] = "cherry"

	fmt.Println("fruits list is: ",fruits)
	fmt.Println("fruits length is: ",len(fruits))

	// 2nd way to create an array
	fruits2 := [4]string{"apple", "banana", "cherry", "grape"}
  fmt.Println("fruits2 list is: ",fruits2)
  fmt.Println("fruits2 length is: ",len(fruits2))

  // 3rd way to create an array
  fruits3 := [...]string{"apple", "banana", "cherry", "grape"}
  fmt.Println("fruits3 list is: ",fruits3)
  fmt.Println("fruits3 length is: ",len(fruits3))

  // 4th way to create an array
  fruits4 := [...]string{}
  fmt.Println("fruits4 list is: ",fruits4)
  fmt.Println("fruits4 length is: ",len(fruits4))

}