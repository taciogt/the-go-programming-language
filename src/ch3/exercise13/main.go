package main

import "fmt"

const (
	KB = 1e3
	MB = KB * 1e3
	GB = MB * 1e3
	TB = GB * 1e3
	PB = TB * 1e3
	EB = PB * 1e3
)

func main() {
	numbers := []int{KB, MB, GB, TB, PB, EB}
	for _, n := range numbers {
		fmt.Printf("%5.3e | %d \n", float64(n), n)
	}
}
