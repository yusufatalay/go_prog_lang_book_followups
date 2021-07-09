// gitcrud is a tool for doing CRUD ops to github issues
package main

import (
	"clio"
	"fmt"
	"github"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func listIssues(tags []string) {
	result, err := github.SearchIssues(tags)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	age := ""     // issues age in strign form
	issueAge := 0 // issues age in terms of hours
	for _, item := range result.Items {
		issueAge = int(time.Now().Sub(item.CreatedAt).Hours())
		if issueAge < 720 {
			age = "Less than a month old"
		} else if issueAge < 8640 {
			age = "Less than a year old"
		} else {
			age = "More than a year old"
		}
		fmt.Printf("#%-5d %9.9s %.55s --->%10s\n", item.Number, item.User.Login, item.Title, age)
	}

}

// rightnow user is expected to enter ownernrepo as user/repo format
func getIssue(ownernrepo, issuenumber string) {
	result, err := github.ReadIssue(ownernrepo, issuenumber)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Issue Number	: %d\n", result.Number)
	fmt.Printf("Created By	: %s\n", result.User.Login)
	fmt.Printf("Created At	: %v\n\n", result.CreatedAt)
	fmt.Printf("\t\t%s\t\n", result.Title)
	fmt.Printf("\n%s\n", result.Body)

}

func createIssue(ownernrepo string) {

	titlebyte, _ := clio.CaptureInputFromEditor(clio.GetPreferredEditorFromEnviroment)
	titletxt := string(titlebyte)

	bodybyte, _ := clio.CaptureInputFromEditor(clio.GetPreferredEditorFromEnviroment)
	bodytxt := string(bodybyte)

	// get the auth key from key.txt file
	keybyte, err := ioutil.ReadFile("key.txt")
	if err != nil {
		log.Fatalf("Error occured while reading the key file : %v", err)
	}
	key := string(keybyte)
	// removing the trailing newline character
	key = strings.TrimSuffix(key, "\n")

	issuenum, err := github.CreateIssue(ownernrepo, titletxt, bodytxt, key)

	if err != nil {
		log.Fatalf("an error has occuded: %v\n", err)
	}

	fmt.Printf("Created issue's number is : %d\n", issuenum)
}

func updateIssue(ownernrepo, issuenumber string) {

	titlebyte, _ := clio.CaptureInputFromEditor(clio.GetPreferredEditorFromEnviroment)
	titletxt := string(titlebyte)

	bodybyte, _ := clio.CaptureInputFromEditor(clio.GetPreferredEditorFromEnviroment)
	bodytxt := string(bodybyte)

	// get the auth key from key.txt file
	keybyte, err := ioutil.ReadFile("key.txt")
	if err != nil {
		log.Fatalf("Error occured while reading the key file : %v", err)
	}
	key := string(keybyte)
	// removing the trailing newline character
	key = strings.TrimSuffix(key, "\n")

	issuenum, err := github.UpdateIssue(ownernrepo, titletxt, bodytxt, key, issuenumber)

	if err != nil {
		log.Fatalf("an error has occuded: %v\n", err)
	}

	fmt.Printf("Updated issue's number is : %d\n", issuenum)
}

func main() {
	updateIssue(os.Args[1], os.Args[2])
}
