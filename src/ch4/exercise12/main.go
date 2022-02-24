package main

import (
	"fmt"
	"go-programming-language/src/ch4/exercise12/xkcd"
)

func main() {
	comicInfo, err := xkcd.GetComicInfo(571)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nComic Info: %+v\n", comicInfo)

	comicInfo, err = xkcd.GetComicInfo(4000)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nComic Info: %+v", comicInfo)

}
