package main

import (
	"bytes"
	"fmt"
)

func main() {
	nums := []int{1, 12, 123, 1234, 12345, 123456, 1234567}
	for _, n := range nums {
		fmt.Printf("%d: %s\n", n, comma(fmt.Sprintf("%d", n)))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	l := len(s)
	for idx, c := range s {
		if idx > 0 && (l-idx)%3 == 0 {
			buf.WriteRune('.')
		}
		buf.WriteRune(c)
	}
	return buf.String()
}
