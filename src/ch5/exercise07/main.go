// go run main.go http://gopl.io
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "outline2: %v\n", err)
		os.Exit(1)
	}

	HtmlPrettyPrinter(resp.Body, os.Stdout)
}

var multilineNodes map[atom.Atom]struct{}

func init() {
	multilineNodes = make(map[atom.Atom]struct{})
	for _, a := range []atom.Atom{atom.Html, atom.Head, atom.Body} {
		multilineNodes[a] = struct{}{}
	}
}

// HtmlPrettyPrinter takes a html document and formats it
func HtmlPrettyPrinter(r io.Reader, w io.Writer) {
	doc, err := html.Parse(r)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement(w), endElement(w))
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre os called before the children are visited (preorder) and
// post is called after (postorder)
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

var depth int

func startElement(w io.Writer) func(*html.Node) {
	return func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Fprintf(w, "%*s<%s>", depth*2, "", n.Data)

			if _, ok := multilineNodes[n.DataAtom]; ok {
				fmt.Fprintf(w, "\n")
			}
			depth++
		} else {
			fmt.Fprintf(w, "%s", n.Data)
		}
	}
}

func endElement(w io.Writer) func(*html.Node) {
	return func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			if _, ok := multilineNodes[n.DataAtom]; ok {
				fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
			} else {
				fmt.Fprintf(w, "</%s>\n", n.Data)
			}
		}
	}
}
