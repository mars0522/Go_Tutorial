package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println(ptr)

	number := 300

	myptr := &number

	fmt.Println(myptr)
	fmt.Println(*myptr)

	*myptr = *myptr + 1
	fmt.Println(*myptr)
}