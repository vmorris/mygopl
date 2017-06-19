// ElementByID traverses an HTML doc tree and stops whenever a match is found
// usage: first arg is url to search, 2nd arg is element ID to search for
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: ElementByID <url> <id>")
		os.Exit(1)
	}
	url := os.Args[1]
	id := os.Args[2]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("ElementByID: failed to get %s\n", url)
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("ElementByID: to parse %s\n", url)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", ElementByID(doc, id))

}

// ElementByID traverses an HTML doc tree and stops whenever a match is found
func ElementByID(doc *html.Node, id string) *html.Node {

	return doc
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
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
