package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
)

func main() {
	hashMethod := flag.String("hash", "sha256", "hash method to be used")
	flag.Parse()

	if len(flag.Args()) == 0 {
		log.Fatalf("missing input argument")
	}
	input := flag.Arg(0)

	var sha []byte
	switch *hashMethod {
	case "sha256":
		a := sha256.Sum256([]byte(input))
		sha = a[:]
	case "sha384":
		a := sha512.Sum384([]byte(input))
		sha = a[:]
	case "sha512":
		a := sha512.Sum512([]byte(input))
		sha = a[:]
	default:
		log.Fatalf("invalid hash method provided: %s", *hashMethod)

	}

	fmt.Printf("SHA256 for \"%s\": %x\n", input, sha)
}
