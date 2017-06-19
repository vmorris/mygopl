// Copyright © 2017 Vance Morris
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// NumElements counts the number of elements of all types given an HTML doc tree
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "numElements: %v\n", err)
		os.Exit(1)
	}
	elementMap := make(map[string]int)
	numElements(elementMap, doc)
	for k, v := range elementMap {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func numElements(m map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		m[n.Data]++
	}
	if n.FirstChild != nil {
		numElements(m, n.FirstChild)
	}
	if n.NextSibling != nil {
		numElements(m, n.NextSibling)
	}
}

/*
type Node struct {
    Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node

    Type      NodeType
    DataAtom  atom.Atom
    Data      string
    Namespace string
    Attr      []Attribute
}
*/
