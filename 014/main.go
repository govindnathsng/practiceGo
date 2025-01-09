package main

import (
	"fmt"
)

// Program: Communication between Goroutines

func calculateSquare(num int, ch chan int) {
	ch <- num * num
}

func main() {
	ch := make(chan int)
	go calculateSquare(4, ch)
	result := <-ch
	fmt.Println("Square is:", result)
}
