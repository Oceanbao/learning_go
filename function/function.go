package main

import (
	"fmt"
)

// function as type in interface
type shape interface {
	area() int
	getType() string
}

type rectangle struct {
	length int
	breath int
}

func (r *rectangle) area() int {
	return r.length * r.breath
}

func (r *rectangle) getType() string {
	return "rectangle"
}

// user defined type
type areaFuncType func(int, int) int

func printArea(x, y int, a areaFuncType) {
	fmt.Printf("Area is: %d\n", a(x, y))
}

func getAreaFunc() areaFuncType {
	return func(x, y int) int {
		return x * y
	}
}

// closure
func getModules() func(int) int {
	count := 0
	return func(x int) int {
		count++
		fmt.Printf("modulus func called %d times\n", count)
		if x < 0 {
			x = x * -1
		}
		return x
	}
}


func main() {

	var shapes []shape
	r := &rectangle{length: 2, breath: 3}
	shapes = append(shapes, r)

	for _, shape := range shapes {
		fmt.Printf("Type: %s, Area %d\n", shape.getType(), shape.area())
	}

	areaF := getAreaFunc()
	printArea(3, 4, areaF)

	modulus := getModules()
	modulus(-1)
	modulus(2)
	modulus(-5)
}
