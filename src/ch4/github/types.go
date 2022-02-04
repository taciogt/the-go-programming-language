// Package github provides a Go API for the Github issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import (
	"time"
)

const (
	baseURL   = "https://api.github.com"
	IssuesURL = baseURL + "/search/issues"
)

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type CreateIssueBody struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Authorization struct {
	User  string
	Token string
}
