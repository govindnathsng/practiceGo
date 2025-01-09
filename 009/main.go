package main

import "fmt"

// Program: Using pointers to modify a value
func main() {
	x := 10
	p := &x

	fmt.Println("Value of x:", *p) // Access the value
	*p = 20                        // Modify the value through pointer
	fmt.Println("Updated value of x:", x)
}
