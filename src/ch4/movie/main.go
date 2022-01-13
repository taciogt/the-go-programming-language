package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	movies := []Movie{
		{Title: "casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Duke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve Mcqueen", "Jaqueline Bisset"}},
	}
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	indentedData, err := json.MarshalIndent(movies, "", "  ")
	fmt.Printf("%s\n", indentedData)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("%#v\n", titles)
	fmt.Println(titles)
}
