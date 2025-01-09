package main

import "github.com/fatih/color"

//Program: Using a third-party package (github.com/fatih/color)
// go mod init demo
// go get github.com/fatih/color

func main() {
	color.Red("This is red text!")
	color.Green("This is green text!")
}
