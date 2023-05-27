package main

import (
	"fmt"
	"time"
)

func main() {

	curr_time := time.Now()
	fmt.Println(curr_time)

	fmt.Println(curr_time.Format("01-02-2006 monday"))
}