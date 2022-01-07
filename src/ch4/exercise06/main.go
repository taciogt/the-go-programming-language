package main

import (
	"fmt"
	"unicode"
)

func main() {
	testStrings := []string{
		"abc",
		"aabac",
		"aa bac",
		"aa ba c",
		"aa ba   c",
		"aa ba  \n c",
		"aa\t\n ba  \n c",
	}

	for _, s := range testStrings {
		byteSlice := []byte(s)
		fmt.Printf("%s: squashSpaces(%v) -> ", s, byteSlice)
		squashSpaces(&byteSlice)
		fmt.Printf("%v\n", byteSlice)
	}
}

func squashSpaces(bs *[]byte) {
	size := len(*bs)

	for i := 0; i < size-1; i++ {
		if unicode.IsSpace(rune((*bs)[i])) && unicode.IsSpace(rune((*bs)[i+1])) {
			copy((*bs)[i:], (*bs)[i+1:])
			i--
			size--
		}
	}
	*bs = (*bs)[:size]
}
