package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		w, i, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "CountWordsAndImages: %v\n", err)
			continue
		}
		fmt.Fprintf(os.Stdout, "%v words, %v images", w, i)
	}
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	visit(&words, &images, n)
	return
}

// visit appends to links each link found in n and returns the result.
func visit(words, images *int, n *html.Node) {
	// count text in text nodes
	if n.Type == html.TextNode {
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			*words++
		}
	} else if n.Type == html.ElementNode && n.Data == "img" {
		*images++
	}
	if !(n.Data == "script" || n.Data == "style") { // skip script and style
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(words, images, c)
		}
	}
	return
}
