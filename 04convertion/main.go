package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Pleaes give us rating between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input,_ := reader.ReadString('\n')

	fmt.Println("Thanks for rating us")
	newRating,err :=  strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil {
		fmt.Println("Oops we got an error while string convertion")
	}else{

		fmt.Println("Your new rating is: ",newRating+1)

	}
	
}