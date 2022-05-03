// Findlinks1 prints the links in an HTML document read from standard input.
// example: go run ../../ch1/fetch https://golang.org | go run main.go
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "exercise 02: %v\n", err)
		os.Exit(1)
	}

	elementsCounter := make(map[string]int)
	dfs(doc, func(n *html.Node) {
		if n.Type == html.ElementNode {

			elementsCounter[n.Data]++
		}
	})
	for name, count := range elementsCounter {
		fmt.Printf("%-7s -> %d\n", name, count)
	}
}

// dfs executes a giver func on each node of a html tree, using a depth first search
func dfs(n *html.Node, f func(n *html.Node)) {
	f(n)

	if n.FirstChild != nil {
		dfs(n.FirstChild, f)
	}
	if n.NextSibling != nil {
		dfs(n.NextSibling, f)
	}
	return
}
