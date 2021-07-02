package main

import (
	"os"
	"strings"
)

func ConcatenateWithSimpleFor(args []string) string {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func ConcatenateWithForInRange(args []string) string {
	var s, sep string
	for _, s := range args[1:] {
		s += sep + s
		sep = " "
	}
	return s
}

func ConcatenateWithJoin(args []string) string {
	return strings.Join(args[1:], " ")
}
