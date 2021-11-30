package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	nums := []int{1, 12, 123, 1234, 12345, 123456, 1234567}
	for _, n := range nums {
		fmt.Printf("%d: %s\n", n, comma(fmt.Sprintf("%d", n)))
	}

	fStrings := []string{"1", "12", "123.45", "1222.3333"}
	for _, n := range fStrings {
		fmt.Printf("%s: %s\n", n, comma(n))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	l := strings.Index(s, ".")
	if l == -1 {
		l = len(s)
	}

	for idx, c := range s[:l] {
		if idx > 0 && (l-idx)%3 == 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune(c)
	}
	buf.WriteString(s[l:])

	return buf.String()
}
