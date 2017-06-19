// Findlinks prints the links in an HTML document read from standard input.
// includes a and style href, and img and script src
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	// anchor and style elements
	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "style") {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	// img and script elements
	if n.Type == html.ElementNode && (n.Data == "img" || n.Data == "script") {
		for _, a := range n.Attr {
			if a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	return links
}
