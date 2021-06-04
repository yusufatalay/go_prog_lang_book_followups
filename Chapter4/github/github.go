// Package github provides a Go API fot the GitHub issue tracker
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const AllIssuesURL = "https://api.github.com/search/issues"
const IssueURL = "https://api.github.com/repos/"

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
	Body      string    // in md format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the Github issue tracker

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(AllIssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// we must close resp.Body on all execution paths
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

// ReadIssue gathers the issue that corresponds the given "issue number"

func ReadIssue(ownernrepo string, issuenumber int) (*Issue, err) {
	q := url.QueryEscape(IssueURL + ownernrepo + "/issues/" + strconv.Itoa(issuenumber))

	resp, err := http.Get(q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Cannot found the issue may be removed or replaced: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil

}
