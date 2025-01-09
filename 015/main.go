package main

import (
	"fmt"
)

// Program 1: Sum an array using Goroutines

func sum(arr []int, ch chan int) {
	total := 0
	for _, v := range arr {
		total += v
	}
	ch <- total
}

func main() {
	arr := []int{1, 2, 8, 4, 5}
	ch := make(chan int)
	go sum(arr, ch)
	fmt.Println("Sum of array:", <-ch)
}
