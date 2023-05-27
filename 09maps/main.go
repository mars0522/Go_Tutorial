package main

import "fmt"

func main() {

	// how to create a map in Golang
	languages := make(map[int]string)
	languages[1] = "Go"
	languages[2] = "Python"
	languages[3] = "Java"
	languages[4] = "Ruby"
	languages[5] = "C"

	fmt.Println(languages)

	// deleting any key from the map
	delete(languages, 1)
	fmt.Println(languages)

	// getting the value from the map
	value, ok := languages[1]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println(ok)
	}

	languages[1] = "Golang"

	// How to Iterate over map in golang
	for key, value := range languages {
		fmt.Printf("for Key %v -> value %v\n", key, value)
	}

}
