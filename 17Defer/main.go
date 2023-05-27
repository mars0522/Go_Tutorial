package main

import "fmt"

func main() {

	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()

}

// expected output -> Hello 43210 two one

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
