package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms":            {"data structures"},
	"calculus":              {"linear algebra"},
	"compilers":             {"data structures", "formal languages", "computer organization"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro do programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(key string)
	visitAll = func(k string) {
		for _, item := range m[k] {
			visitAll(item)
			if !seen[item] {
				seen[item] = true
				visitAll(item)
				order = append(order, item)
			}
		}
		if !seen[k] {
			seen[k] = true
			order = append(order, k)
		}
	}
	for k := range m {
		visitAll(k)
	}

	return order
}
