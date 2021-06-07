// Package github provides a Go API fot the GitHub issue tracker
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
	"log"
	"bytes"
	"errors"
)
// AllIssuesURL is being used for listing all issues with given tags 
const AllIssuesURL = "https://api.github.com/search/issues"
// IssueURL is for editing, creating, updating and reading spesific issues
const IssueURL = "https://api.github.com/repos/"
// IssesSearchResult holds a list of issue 
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}
// Issue struct holds generic fields 
type Issue struct {
	Number    int	`json:"number"`
	HTMLURL   string `json:"html_url"`
	Title     string `json:"title"`
	State     string `json:"state"`
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string `json:"body"` // in md format
}
// User struct holds its field from the JSON gthered from any response body
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
func ReadIssue(ownernrepo, issuenumber string) (*Issue, error) {

	resp, err := http.Get(IssueURL+ownernrepo + "/issues/" + issuenumber)
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


// CreateIssue creates a new issue on given repo with given title and body(body is optional) 
func CreateIssue(ownernrepo, title, body, authtoken string) (int , error){

	// check if user provides an authorization key if not then throw an error
	if authtoken == ""{
		errNoAuth := errors.New("[!]Authorization keys not found please put you key into a authkey.txt file")
		return 0, errNoAuth
	}

	// create an issue struct instance
	issuetosend := Issue{Title : title, Body: body}
	// then convert int to json in order to send it with POST request
	issuejson, err := json.Marshal(issuetosend)

	if err != nil {
		log.Fatalf("%v",err)
		return 0, err
	}
	// encode the data
	responsebody := bytes.NewBuffer(issuejson)
	// constructing the target URL 
	reqURL := IssueURL+ ownernrepo+ "/issues"
	// in order to add custom headers to any request we have to create a client
	client := &http.Client{}

	req, _ := http.NewRequest("POST",reqURL,responsebody)
	// setting the headers as recommended
	req.Header.Set("Authorization","token "+authtoken)
	req.Header.Set("Accept","application/vnd.github.v3+json")

	// make a post request 
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("%v",err)
		return 0, err
	}

	//  close the response body after everything is done
	defer resp.Body.Close()



	statcode := resp.StatusCode
	if statcode == 201{
		fmt.Println("Issue created successfully")
		var result Issue
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
		}
		return result.Number, nil
	}else if statcode == 403{
		errForbidden := errors.New("403-Forbidden")
		return 0 , errForbidden
	}else if statcode == 404{
		errNotFound := errors.New("404-Not Found")
		fmt.Println(IssueURL+ ownernrepo+ "/issues")
		return 0 , errNotFound
	}else if statcode == 410{
		errGone := errors.New("410-Gone")
		return 0 , errGone
	}else if statcode == 422{
		errNotValid := errors.New("422-Validation Failed Unprocessable Entity")
		return 0 , errNotValid
	}else  {
		errUnavailable := errors.New("503-Service Unavailable")
		return 0 , errUnavailable
	}
}

// UpdateIssue updates given issue with provided issue number
func UpdateIssue(ownernrepo, title, body, authtoken, issuenumber string) (int , error){
// this function looks identical with the CreateIssue function just the http method and URL is different
	// check if user provides an authorization key if not then throw an error
	if authtoken == ""{
		errNoAuth := errors.New("[!]Authorization keys not found please put you key into a key.txt file")
		return 0, errNoAuth
	}

	// create an issue struct instance 
	issuetosend := Issue{Title : title, Body: body}
	// then convert int to json in order to send it with POST request
	issuejson, err := json.Marshal(issuetosend)

	if err != nil {
		log.Fatalf("%v",err)
		return 0, err
	}
	// encode the data
	responsebody := bytes.NewBuffer(issuejson)
	// constructing the target URL 
	reqURL := IssueURL+ ownernrepo+ "/issues/"+issuenumber
	// in order to add custom headers to any request we have to create a client
	client := &http.Client{}

	req, _ := http.NewRequest("PATCH",reqURL,responsebody)
	// setting the headers as recommended
	req.Header.Set("Authorization","token "+authtoken)
	req.Header.Set("Accept","application/vnd.github.v3+json")

	// make a post request 
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("%v",err)
		return 0, err
	}

	//  close the response body after everything is done
	defer resp.Body.Close()



	statcode := resp.StatusCode
	if statcode == 200{
		fmt.Println("Issue updated successfully")
		var result Issue
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
		}
		return result.Number, nil
	}else if statcode == 301{
		errMovedPerm := errors.New("301- MOved Permanently")
		return 0 , errMovedPerm
	}else if statcode == 403{
		errForbidden := errors.New("403-Forbidden")
		return 0 , errForbidden
	}else if statcode == 404{
		errNotFound := errors.New("404-Not Found")
		fmt.Println(IssueURL+ ownernrepo+ "/issues")
		return 0 , errNotFound
	}else if statcode == 410{
		errGone := errors.New("410-Gone")
		return 0 , errGone
	}else if statcode == 422{
		errNotValid := errors.New("422-Validation Failed Unprocessable Entity")
		return 0 , errNotValid
	}else  {
		errUnavailable := errors.New("503-Service Unavailable")
		return 0 , errUnavailable
	}
}
