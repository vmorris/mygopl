// Findlinks3 crawls the web, starting with the URLs on the command line.

// TODO: this is not complete -- only detecting when a found link is not
// part of the original domain. there' also a panic index out of range
// to handle properly...
package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"gopl.io/ch5/links"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func save(u *url.URL) error {

	domain := getDomain(u)

	if startDomain == "" {
		startDomain = domain
	} else if domain != startDomain {
		return fmt.Errorf("unable to save site %v: outside domain", domain)
	} else if len(domain) == 0 || !strings.Contains(domain, ".") {
		return fmt.Errorf("unable to save site %v: not valid", domain)
	}

	return nil
}

func crawl(u string) []string {
	fmt.Println(u)

	_u, err := url.Parse(u)
	if err != nil {
		fmt.Printf("crawl: unable to parse %v: %v:", u, err)
	}

	err = save(_u)
	if err != nil {
		fmt.Println(err)
	}

	list, err := links.Extract(u)
	if err != nil {
		fmt.Printf("crawl: unable to extract links from %v: %v:", u, err)
	}
	return list
}

func getDomain(u *url.URL) string {
	components := strings.Split(u.Hostname(), ".")
	domain := components[len(components)-2] + "." + components[len(components)-1]
	return domain
}

var startDomain string

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	url := os.Args[1:]
	breadthFirst(crawl, url)
}
