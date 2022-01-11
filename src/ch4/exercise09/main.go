// go run main.go --file test.txt
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := flag.String("file", "", "path for file to be examined")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	wf := wordfreq(f)

	for word, count := range wf {
		fmt.Printf("%d\t%s\n", count, word)
	}
}

// wordfreq counts the frequency of each word in a text file
func wordfreq(file *os.File) map[string]int {
	result := make(map[string]int)

	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		result[word]++
	}
	return result
}
