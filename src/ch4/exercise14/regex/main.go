package main

import (
	"fmt"
	"regexp"
)

func main() {

	path := "/issues/1234"

	issueID := regexp.MustCompile(`/issues/(\d+)`)
	fmt.Printf("FindAll(): %s\n", issueID.FindAll([]byte(path), -1))
	fmt.Printf("FindSubmatch(): %s\n", issueID.FindSubmatch([]byte(path)))
}
