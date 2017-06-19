// Copyright © 2017 Vance Morris
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// ShowText traverses an HTML doc and prints all the contents of text nodes

// this solution isn't complete -- it needs to not print script or style
// contents. Potentially we can pass a flag down into the recursion, but I'm
// not too keen on it. Also, we could not print any strings with "function",
// for example, but this could certainly hit normal HTML data strints too!
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func showText(n *html.Node) {
	if n.Type == html.TextNode {
		if len(strings.TrimSpace(n.Data)) > 0 {
			fmt.Println(n.Data)
		}
	}
	if !(n.Data == "script" || n.Data == "style") { // skip script and style
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			showText(c)
		}
	}
}

func main() {

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "showText: %v\n", err)
		os.Exit(1)
	}
	showText(doc)
}
