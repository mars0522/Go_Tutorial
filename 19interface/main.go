package main

import (
	"fmt"
	"math"
)

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}
func (r *circle) area() float64 {
	return math.Pi * r.radius * r.radius
}

type shapes interface {
	area() float64
}

func main() {

	r := rect{width: 10, height: 10}
	c := circle{radius: 10}

	objects := []shapes{&r, &c}

	for _, obj := range objects {
		fmt.Println(obj.area())
	}

}
