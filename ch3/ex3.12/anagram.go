// Anagram reports whether two strings given as CLI args are anagrams of each other
//
// Example:
//	$ ./anagram asdf sadf
//  true
//	$ ./anagram cathat asdf
//  false
//
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	usage := "usage: ./anagram string1 string2"
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	string1, string2 := os.Args[1], os.Args[2]

	if len(string1) != len(string2) {
		fmt.Println("false")
		os.Exit(0)
	}

	for _, c := range string1 {
		if ip := strings.IndexRune(string2, c); ip >= 0 {
			string2 = string2[:ip] + string2[ip+1:]
		}
	}

	if len(string2) == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
