package main

import "fmt"

func main() {
	testStrings := [][]string{
		{"a", "b", "c"},
		{"a", "a", "b", "a", "c"},
	}

	for _, ss := range testStrings {
		fmt.Printf("eliminateAdj(%v) -> ", ss)
		eliminateAdjacentDuplicates(&ss)
		fmt.Printf("%v\n", ss)
	}
}

func eliminateAdjacentDuplicates(ssPtr *[]string) {
	ss := *ssPtr
	size := len(ss)
	for i := 0; i < size-1; i++ {
		if ss[i] == ss[i+1] {
			copy(ss[i:], ss[i+1:])
			i--
			size--
		}
	}
	*ssPtr = ss[:size]
}
