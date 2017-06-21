package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// ElementsByTagName parses an html.Node and returns all child nodes
// that match the tags named
func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {

	results := []*html.Node{}

	return results

}

func outline(url string) error {

	var depth int

	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}

	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func main() {

	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	ElementsByTagName(doc, os.Args[2:]...)

}
