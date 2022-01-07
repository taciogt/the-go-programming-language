package main

import "fmt"

func main() {
	words := [][]byte{
		[]byte("abcdef"),
	}
	for _, w := range words {
		originalWord := string(w)
		reverse(w)
		fmt.Printf("reverse(%s) -> %s", originalWord, w)
	}
}

func reverse(word []byte) {
	size := len(word)
	for i := 0; i < size/2; i++ {
		j := size - 1 - i
		word[i], word[j] = word[j], word[i]
	}
}
