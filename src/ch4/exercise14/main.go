package main

import (
	"github.com/taciogt/go-programming-language/src/ch4/exercise14/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/issues/", handlers.GetIssuesHandler)
	http.HandleFunc("/", handlers.ListIssuesHandler)

	log.Println("starting http server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
