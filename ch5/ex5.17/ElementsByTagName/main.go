package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// ElementsByTagName parses an html.Node and returns all child nodes
// that match the tags named
func ElementsByTagName(doc *html.Node, tags ...string) []html.Node {
	results := []html.Node{}
	checkTag := func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, t := range tags {
				if n.Data == t {
					results = append(results, *n)
				}
			}
		}
	}
	forEachNode(doc, checkTag, nil)
	return results
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

	results := ElementsByTagName(doc, os.Args[2:]...)
	for _, f := range results {
		fmt.Printf("%v\n", f)
	}
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
