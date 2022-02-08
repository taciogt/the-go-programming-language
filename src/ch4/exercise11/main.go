// Script to interact with the issues of any github repository
// go run main.go -repo taciogt/test-github-api
// go run main.go -help
// create an issue
// go run main.go -repo taciogt/test-github-api -create -text "some issue content"
// go run main.go -repo taciogt/test-github-api -create -title "New Test Title 4" -body "Issue description"

package main

import (
	"flag"
	"fmt"
	"go-programming-language/src/ch4/github"
	"log"
	"os"
)

func main() {
	var repository string
	flag.StringVar(&repository, "repo", "", "github repository to interact with (taciogt/test-github-api)")

	var create bool
	flag.BoolVar(&create, "create", false, "creates an issue on github")
	var update bool
	flag.BoolVar(&update, "update", false, "updates an issue on github")
	var list bool
	flag.BoolVar(&list, "list", false, "lists issues of a given repository")

	var title string
	flag.StringVar(&title, "title", "", "title to be used when creating the issue")
	var body string
	flag.StringVar(&body, "body", "", "description for the issue")
	var issueNumber int
	flag.IntVar(&issueNumber, "number", 0, "issue number to be altered")

	flag.Parse()

	authorization := github.Authorization{
		User:  os.Getenv("GITHUB_USER"),
		Token: os.Getenv("GITHUB_ACCESS_TOKEN"),
	}

	log.Printf("repository: %s", repository)

	if create {
		log.Println("creating issue operation")
		if title == "" {
			log.Fatal("title must be provided for create operation")
		}
		log.Printf("title: %s\n", title)

		if body == "" {
			log.Fatal("body must be provided for create operation")
		}

		err := github.CreateIssue(authorization, repository, title, body)
		if err != nil {
			log.Fatal(err)
		}
	} else if update {
		log.Println("updating issue operation")
		if issueNumber == 0 {
			log.Fatal("issue number must be defined and different than 0")
		}

		err := github.UpdateIssue(authorization, repository, issueNumber, title, body)
		if err != nil {
			log.Fatal(err)
		}
	} else if list {
		log.Printf("listing issues for repository %s\n", repository)

		issues, err := github.ListIssues(repository)
		if err != nil {
			log.Fatal(err)
		}

		for _, issue := range issues {
			fmt.Printf("%+v\n", issue)
		}
	} else {
		log.Fatal("ERROR: no operation provided")
	}
}
