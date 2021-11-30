package main

import (
	"fmt"
	"reflect"
)

func main() {
	stringPairs := [][]string{
		{"abc", "cba"},
		{"abba", "baba"},
		{"abc", "ab"},
		{"abba", "aba"},
	}
	for _, p := range stringPairs {
		fmt.Printf("%s, %s is anagrams? %t\n", p[0], p[1], isAnagram(p[0], p[1]))
	}
}

func isAnagram(s1 string, s2 string) bool {
	runeCount1 := make(map[rune]int)
	runeCount2 := make(map[rune]int)
	for _, r := range s1 {
		runeCount1[r]++
	}
	for _, r := range s2 {
		runeCount2[r]++
	}
	return reflect.DeepEqual(runeCount1, runeCount2)
}
