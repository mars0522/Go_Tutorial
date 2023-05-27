package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitList = []string{"apple", "banana"}
	fmt.Println(fruitList)
	fmt.Printf("Type of FruitList is %T \n",fruitList)

	fruitList = append(fruitList, "orange","mango")
	fmt.Println(fruitList)

	fruitList1 := append(fruitList[1:3])
	fmt.Println("FruitList 1 is : ",fruitList1)
	fmt.Printf("Type of FruitList 1 is %T \n",fruitList1)
	fmt.Println("FruitList 2 is : ",fruitList[2:])


	// now creating a slice using make() method

	scores := make([]int, 2)
	scores[0] = 10
	scores[1] = 20

	fmt.Println(scores)
	fmt.Println(len(scores))

	scores = append(scores, 345,323,342)
	fmt.Println(scores)
	fmt.Println(len(scores))

	sort.Ints(scores)
	fmt.Println(scores)

	// how to remove the values from the slice based on the indexes
	numbers := []int{1,2,3,4,5,6,7,8,9}
	
	index := 4

	numbers = append(numbers[:index],numbers[index+1:]...)
	fmt.Println("After deleting the number at index 4 slice looks like : ",numbers)

	


}