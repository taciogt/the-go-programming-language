// go run main.go http://gopl.io
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "outline2: %v\n", err)
		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	elm := ElementByID(doc, os.Args[2])
	fmt.Println("element found:", elm)
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var found html.Node

	forEachNode(doc, checkNodeForID(&found, id), nil)
	return &found
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre os called before the children are visited (preorder) and
// post is called after (postorder)
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil && pre(n) {
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil && post(n) {
		return
	}
}

func checkNodeForID(found *html.Node, id string) func(*html.Node) bool {
	return func(n *html.Node) bool {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				*found = *n
				return true
			}
		}
		return false
	}
}
