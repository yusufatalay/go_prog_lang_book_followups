// gitCRUD is a tool for doing CRUD ops to github issues
package main

import (
	"fmt"
	"github"
	"log"
	"os"
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

func main() {
	getIssue(os.Args[1], os.Args[2])
}
