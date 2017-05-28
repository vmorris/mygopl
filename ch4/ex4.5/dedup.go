package main

import "fmt"

func main() {
	s := []string{"a", "a", "b", "c", "d", "d", "d", "e"}
	s = dedup(s)
	fmt.Println(s)
}

func dedup(str []string) []string {
	out := str[0:1] // always pick up the first element
	p := 0          // pointer into out
	for _, s := range str[1:] {
		if s == out[p] {
			continue // dup found, skip
		} else {
			out = append(out, s)
			p++
		}
	}
	return out
}
