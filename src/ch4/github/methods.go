// https://docs.github.com/en/rest/overview/other-authentication-methods#via-oauth-and-personal-access-tokens
// https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token
package github

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// SearchIssues queries the Github issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

// CreateIssue creates an issue on a given Github repository
// Github api: https://docs.github.com/en/rest/reference/issues#create-an-issue
func CreateIssue(authorization Authorization, repo string, title string, description string) error {
	client := &http.Client{}

	createIssue := CreateIssueBody{Title: title, Body: description}
	body, err := json.Marshal(createIssue)
	if err != nil {
		return err
	}

	url_ := fmt.Sprintf("%s/repos/%s/issues", baseURL, repo)
	req, err := http.NewRequest("POST", url_, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "Basic "+basicAuth(authorization.User, authorization.Token))

	log.Printf("create issue request: url = %s, description = %v", req.URL, req.Body)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	log.Printf("create issue response: status = %s", resp.Status)

	return nil
}

// ListIssues list issues of given Github repository
// Github api: https://docs.github.com/en/rest/reference/issues#list-repository-issues
func ListIssues(repo string) ([]Issue, error) {
	client := &http.Client{}
	url_ := fmt.Sprintf("%s/repos/%s/issues", baseURL, repo)
	req, err := http.NewRequest("GET", url_, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []Issue
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Printf("response status = %s", resp.Status)
		return nil, err
	}

	return result, nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
