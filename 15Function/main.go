package main

import "fmt"

func main() {

	res, msg := adder(1, 2, 3, 4, 5, 6)

	fmt.Println(msg, res)

}

func adder(values ...int) (int, string) {
	sum := 0

	for _, value := range values {
		sum += value
	}
	return sum, "Your sum is"
}
