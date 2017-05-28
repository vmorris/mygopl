package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	rotate(s, 3)
	fmt.Println(s)
}

func rotate(s []int, n int) {
	length := len(s)
	temp := make([]int, n)
	copy(temp, s[:n])
	copy(s, s[n:])
	copy(s[length-n:], temp)
}
