// ElementByID traverses an HTML doc tree and stops whenever a match is found
// usage: first arg is url to search, 2nd arg is element ID to search for

//NOTE: totally cheated here, pre and post returning bools was a mess!

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
	n, found := forEachNode(doc, id)
	if found {
		return n
	}
	return nil
}

func forEachNode(n *html.Node, id string) (*html.Node, bool) {
	for _, a := range n.Attr {
		if a.Key == "id" && a.Val == id {
			//	fmt.Fprintf(os.Stdout, "FOUND id=%v\n", a.Val)
			//	fmt.Fprintf(os.Stdout, "%v\n", n)
			return n, true
		}
	}

	node := new(html.Node)
	var found bool

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node, found = forEachNode(c, id)
		if found {
			break
		}
	}
	return node, found
}
