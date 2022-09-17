package main

import (
	"fmt"
	"github.com/taciogt/go-programming-language/src/ch4/exercise13/omdbapi"
	"os"
)

func main() {
	apiKey := os.Getenv("OPEN_MOVIE_DB_KEY")
	client := omdbapi.NewClient(apiKey)
	movies, err := client.SearchMovie("alien")
	if err != nil {
		panic(err)
	}

	fmt.Printf("movies found (%d):\n", len(movies))
	for i, m := range movies {
		fmt.Printf("movie #%d: %+v\n", i, m)
		filename, err := client.DownloadPoster(m)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\tdowloaded poster: %s\n", filename)
	}
}
