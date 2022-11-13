package jira

import (
	"io"
	"strings"

	jr "github.com/andygrunwald/go-jira"
	"github.com/labstack/gommon/log"
)

// NewClient returns a new Jira API client.
// A baseURL is a URL in the form of "https://xxx.atlassian.net".
func NewClient(baseURL, email, token string) (*jr.Client, error) {
	tp := jr.BasicAuthTransport{
		Username: email,
		Password: token,
	}
	return jr.NewClient(tp.Client(), baseURL)
}

func CreateSimpleIssue(cli *jr.Client, projectKey, summary, description, issueType string) (*jr.Issue, error) {
	newIssue := jr.Issue{
		Fields: &jr.IssueFields{
			Summary:     summary,
			Description: description,
			Type: jr.IssueType{
				Name: issueType,
			},
			Project: jr.Project{Key: projectKey},
		},
	}
	created, resp, err := cli.Issue.Create(&newIssue)
	if err != nil {
		sb := new(strings.Builder)
		sb.Grow(int(resp.ContentLength))
		io.Copy(sb, resp.Body)
		log.Error(sb.String())
		return nil, err
	}
	return created, nil
}
