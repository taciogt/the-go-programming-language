package templates

import (
	"bytes"
	"go-programming-language/src/ch4/github"
	"html/template"
)

type listIssuesTemplateData struct {
	RepositoryName string
	Issues         []github.Issue
}

func ListIssuesTemplate(issues []github.Issue) ([]byte, error) {
	t, err := template.ParseFiles("templates/list-issues.html")
	if err != nil {
		return nil, err
	}
	data := listIssuesTemplateData{
		RepositoryName: "golang",
		Issues:         issues,
	}
	buf := bytes.Buffer{}
	if err := t.Execute(&buf, data); err != nil {
		return nil, err
	}
	return []byte(buf.String()), nil
}

func GetIssueTemplate(issue github.Issue) ([]byte, error) {
	t, err := template.ParseFiles("templates/get-issue.html")
	if err != nil {
		return nil, err
	}
	data := map[string]interface{}{
		"Issue": issue,
	}
	buf := bytes.Buffer{}
	if err := t.Execute(&buf, data); err != nil {
		return nil, err
	}
	return []byte(buf.String()), nil
}
