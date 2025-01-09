package main

import "fmt"

// Program 1: Function to calculate factorial

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Program 2: Method with a custom type

type Rectangle struct {
	length, breadth float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

func main() {
	// Program 1:
	// fmt.Println("Factorial of 5 is:", factorial(5))
	// Program 2:
	rect := Rectangle{length: 5, breadth: 3}
	fmt.Println("Area of rectangle:", rect.Area())
}
