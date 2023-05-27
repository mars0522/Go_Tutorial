package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7}

	// 1st way of running the loop
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Println("==============================================")
	// 2nd way of running the loop
	for i := range arr {
		fmt.Println(arr[i])
	}
	fmt.Println("==============================================")

	// 3rd way of running the loop
	for index, values := range arr {
		fmt.Printf("Key :%v and value :%v\n", index, values)
	}

}
