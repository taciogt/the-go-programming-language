// Echo4 prints its command-line arguments.
// go build echo4.go
// ./echo4 -s / This is the echo message
// ./echo4 -n -s / This is the echo message

package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
