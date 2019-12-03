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

	var contents []byte
	_, err = f.Read(contents)
	if err != nil {
		fmt.Println("Failed to read file")
	}

	fmt.Println(contents)
}

func anotherFunc() {

	fmt.Println("This is a funciton")
}
