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

	//listComics()
	//comics, err := xkcd.ListAllComics()
	//bw := bufio.NewWriter(os.Stdout)
	//for comic := range xkcd.ListAllComics() {
	//	if _, err := fmt.Fprintf(bw, "comic #%d: %s\n", comic.Num, comic.Title); err != nil {
	//		panic(err)
	//	}
	//	//fmt.
	//}
	//if err := bw.Flush(); err != nil {
	//	panic(err)
	//}
	if err := xkcd.BuildIndex(); err != nil {
		panic(err)
	}

	//comicInfo, err = xkcd.GetComicInfo(4000)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("\nComic Info: %+v", comicInfo)

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
