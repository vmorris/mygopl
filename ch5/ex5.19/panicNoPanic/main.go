package main

import "fmt"

func main() {
	fmt.Printf("%v\n", panicNoPanic())
}

// panicNoPanic has no return statement but returns a non-zero value
// through a panic and recovery mechanisms
func panicNoPanic() (result int) {
	defer dealWithIt(&result)
	panic("but not really")
}

func dealWithIt(r *int) {
	recover()
	*r = 1
}
