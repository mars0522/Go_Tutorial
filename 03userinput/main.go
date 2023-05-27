package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// create the reader and from where it should read
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please provide your rating: ")

	// now reader starts reading but till what , till new line encounters
	input, err := reader.ReadString('\n')

	// if there is no error in readng, print what it has readed
	if err == nil {
		fmt.Println("You have entered the rating of ",input)
	}
}