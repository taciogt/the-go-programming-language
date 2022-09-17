package main

import (
	"fmt"
	"github.com/taciogt/go-programming-language/src/ch2/exercise01/tempconv"
	"github.com/taciogt/go-programming-language/src/ch2/exercise02/lenconv"
	"os"
	"strconv"
)

func convert(n float64) {
	fmt.Printf("Starting conversions of number %v:\n", n)

	f := tempconv.Fahrenheit(n)
	c := tempconv.Celsius(n)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))

	m := lenconv.Meters(n)
	feet := lenconv.Feet(n)
	fmt.Printf("%s = %s, %s = %s\n", m, lenconv.MToF(m), feet, lenconv.FToM(feet))

	fmt.Println("--------")

}

func main() {
	var n float64
	var numbers []float64
	for _, arg := range os.Args[1:] {
		n, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		numbers = append(numbers, n)
	}

	if len(os.Args) == 1 {
		fmt.Print("Provide a number: ")
		_, err := fmt.Scanf("%f", &n)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
		}
		numbers = append(numbers, n)
	}

	for _, n := range numbers {
		convert(n)
	}
}
