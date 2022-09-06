// Exercise 05
// example: go run main.go https://pkg.go.dev/golang.org/x/net/html
package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "CountWordsAndImages(): %v\n", err)
			continue
		}

		fmt.Println(url)
		fmt.Println("words", words)
		fmt.Println("images", images)
		fmt.Println("----------------")
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
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		if n.Data == "img" {
			images++
		}
	} else if n.Type == html.TextNode {
		input := bufio.NewScanner(strings.NewReader(n.Data))
		input.Split(bufio.ScanWords)
		for input.Scan() {
			words++
		}
	}

	childWords, childImages := countWordsAndImages(n.FirstChild)
	words += childWords
	images += childImages

	siblingWords, siblingImages := countWordsAndImages(n.NextSibling)
	words += siblingWords
	images += siblingImages

	return
}
