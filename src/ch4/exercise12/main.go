package main

import (
	"flag"
	"fmt"
	"go-programming-language/src/ch4/exercise12/xkcd"
)

func main() {
	var buildIndex bool
	flag.BoolVar(&buildIndex, "build", false, "build a local index with all available xkcd comics")

	var searchTerm string
	flag.StringVar(&searchTerm, "search", "", "search for the argument on the ")

	flag.Parse()

	if buildIndex {
		if err := xkcd.BuildIndex(); err != nil {
			panic(err)
		}
	}

	if searchTerm != "" {
		comics, err := xkcd.LoadIndex()
		if err != nil {
			panic(err)
		}

		comics, err = xkcd.SearchTerm(comics, searchTerm)
		if err != nil {
			panic(err)
		}

		fmt.Printf("-- comics found: --\n")
		for _, c := range comics {
			fmt.Printf("\nComic Info: %+v\n", c)
		}

	}

}

// TODO: build index using parallel go functions
//func getComicInfo(c chan xkcd.ComicInfo, comicNumber int) {
//	comicInfo, err := xkcd.GetComicInfo(comicNumber)
//	if err != nil {
//		panic(err)
//	}
//	c <- comicInfo
//}
//
//
//func listComics() {
//	c := make(chan xkcd.ComicInfo)
//	for i := 1; i < 25; i++ {
//		go getComicInfo(c, i)
//	}
//
//	for comicInfo := range c {
//		fmt.Printf("comic #%d: %s\n", comicInfo.Num, comicInfo.Title)
//	}
//}
