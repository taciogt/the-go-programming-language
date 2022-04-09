package templates

import (
	"bytes"
	"go-programming-language/src/ch4/github"
	"html/template"
	"log"
	"os"
)

type homeTemplateData struct {
	RepositoryName string
	Issues         []github.Issue
}

func GetHomeTemplate(issues []github.Issue) ([]byte, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	log.Print("working dir:", wd)

	//bs, err := os.ReadFile("templates/home.html")
	//if err != nil {
	//	return nil, err
	//}

	t, err := template.ParseFiles("templates/home.html")
	if err != nil {
		return nil, err
	}
	data := homeTemplateData{
		RepositoryName: "golang",
		Issues:         issues,
	}
	buf := bytes.Buffer{}
	if err := t.Execute(&buf, data); err != nil {
		return nil, err
	}
	return []byte(buf.String()), nil

	//return bs, nil
}
