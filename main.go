package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create file")
	}

	_, err = f.WriteString("Some words")
	if err != nil {
		fmt.Println("Failed to write to file")
	}

	var b []byte
	_, err = f.Read(b)
	if err != nil {
		fmt.Println("Error")
	}
}

func anotherFunc() {

	fmt.Println("This is a funciton")
}
