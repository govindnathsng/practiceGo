package main

import (
	"fmt"
	"time"
)

//Program: Demonstrating concurrency with Goroutines

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printNumbers()
	fmt.Println("Concurrent execution!")
	time.Sleep(3 * time.Second) // Wait for Goroutine to finish
}
