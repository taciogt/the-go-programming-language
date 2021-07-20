// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	filesDuplicatesCount := make(map[string]int)
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, filesDuplicatesCount)
		f.Close()
	}

	fmt.Println("Files with any duplicate line")
	for filename, _ := range filesDuplicatesCount {
		fmt.Printf("- %s\n", filename)
	}
}

func countLines(f *os.File, filesDuplicatesCount map[string]int) {
	counts := make(map[string]int)

	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			filesDuplicatesCount[f.Name()]++
		}
	}
}
