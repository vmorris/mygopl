// ExpandString replaces each found substring with the text returned by another
// function

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: ExpandString <sourceString>")
		os.Exit(1)
	}
	fmt.Println(expand(os.Args[1], foobar))
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}

func foobar(s string) string {
	return "bar"
}
