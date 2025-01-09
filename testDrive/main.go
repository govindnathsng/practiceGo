package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	firstName string
}

func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.firstName))
}

func main() {
	// fmt.Println("Hello!")
	// defer check()
	// test()
	// fmt.Println("midd")
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	// s := []byte("lets check!")

	// _, err = f.Write(s)
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }

	p := person{
		firstName: "Jonny",
	}

	var b bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b)

	fmt.Println(b.String())

}
func check() {
	fmt.Println("first")

}

func test() {
	fmt.Println("second")
}
