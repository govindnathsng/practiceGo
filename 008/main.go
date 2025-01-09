package main

import "fmt"

// Program 1: Check if a number is even or odd
func main() {
	var num int
	fmt.Println("Enter a number:")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}
}
