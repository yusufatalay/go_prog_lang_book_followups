// issues prints a table of github issues matching the search terms
package main

import (
	"fmt"
	"github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
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
