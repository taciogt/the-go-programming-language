package main

import (
	"fmt"
	"os"
)

func main(){
	for i, s := range os.Args[1:] {
		fmt.Printf("%d: %v\n", i, s)
	}
}