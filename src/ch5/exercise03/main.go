// Exercise 03 prints the content of all textg elements
// example: go run ../../ch1/fetch https://golang.org | go run main.go
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "exercise 03: %v\n", err)
		os.Exit(1)
	}

	dfs(doc, func(n *html.Node) {
		if n.Type == html.TextNode {
			trimmedContent := strings.TrimSpace(n.Data)
			if len(trimmedContent) > 0 && n.Parent.Data != "script" && n.Parent.Data != "noscript" {
				fmt.Println("----------------", n.Type, n.Parent.Data)
				fmt.Println(trimmedContent)
			}
		}
	})
	//for name, count := range elementsCounter {
	//	fmt.Printf("%-7s -> %d\n", name, count)
	//}
}

// dfs executes a giver func on each node of a html tree, using a depth first search
func dfs(n *html.Node, f func(n *html.Node)) {
	//fmt.Println("******\n", n.Data, "\n******")
	f(n)

	if n.FirstChild != nil {
		dfs(n.FirstChild, f)
	}
	if n.NextSibling != nil {
		dfs(n.NextSibling, f)
	}
	return
}
