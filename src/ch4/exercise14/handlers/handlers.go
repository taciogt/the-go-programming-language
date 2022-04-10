package handlers

import (
	"go-programming-language/src/ch4/exercise14/templates"
	"go-programming-language/src/ch4/github"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

const githubRepository = "golang/go"

func ListIssuesHandler(w http.ResponseWriter, req *http.Request) {
	log.Print("on listing issues handler")

	issues, err := github.ListIssues(githubRepository)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte(err.Error())); err != nil {
			panic(err)
		}
	}
	//for i, issue := range issues {
	//	fmt.Printf("issue #%d: %+v\n", i, issue)
	//}

	content, err := templates.ListIssuesTemplate(issues)
	if err != nil {
		panic(err)
	}
	if _, err = w.Write(content); err != nil {
		panic(err)
	}
}

func GetIssuesHandler(w http.ResponseWriter, req *http.Request) {
	log.Print("on get issue handler")

	issueIDRegEx := regexp.MustCompile(`/issues/(\d+)`)
	matches := issueIDRegEx.FindSubmatch([]byte(req.URL.String()))
	if len(matches) < 2 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	issueNumber, err := strconv.Atoi(string(matches[1]))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	issue, err := github.GetIssue(githubRepository, issueNumber)
	if err != nil {
		log.Println("error", err)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			log.Fatal(err)
		}

	}

	content, err := templates.GetIssueTemplate(issue)
	if err != nil {
		panic(err)
	}
	if _, err = w.Write(content); err != nil {
		panic(err)
	}

}
