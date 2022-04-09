package main

import (
	"fmt"
	"go-programming-language/src/ch4/exercise14/templates"
	"go-programming-language/src/ch4/github"
	"log"
	"net/http"
)

const repoQuery = "repo:golang/go"
const githubRepository = "golang/go"

func main() {
	//homeHandler := func(w http.ResponseWriter, req *http.Request) {
	//	_, _ = io.WriteString(w, "welcome!\n")
	//}
	//http.HandleFunc("/", homeHandler)
	http.HandleFunc("/", issuesHandler)

	log.Println("starting http server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func issuesHandler(resp http.ResponseWriter, req *http.Request) {
	issues, err := github.ListIssues(githubRepository)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		if _, err = resp.Write([]byte(err.Error())); err != nil {
			panic(err)
		}
	}
	for i, issue := range issues {
		fmt.Printf("issue #%d: %+v\n", i, issue)
	}

	content, err := templates.GetHomeTemplate(issues)
	if err != nil {
		panic(err)
	}
	if _, err = resp.Write(content); err != nil {
		panic(err)
	}

}
