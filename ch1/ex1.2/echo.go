package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Program: " + os.Args[0])
	fmt.Println("Args:")
	for i, v := range os.Args[1:] {
		fmt.Println("Index: " + strconv.Itoa(i) + "\t Value: " + v)
	}
}
