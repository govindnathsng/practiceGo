package main

import "fmt"

func main() {

	// Program 1: Using slices

	numbers := []int{1, 2, 3, 4}
	numbers = append(numbers, 5)

	fmt.Println("Numbers:", numbers)

	// Program 2: Using maps

	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}

	ages["Charlie"] = 35
	fmt.Println("Ages:", ages)
}
