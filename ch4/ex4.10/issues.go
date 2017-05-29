// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/vmorris/mygopl/ch4/ex4.10/github"
)

//!+
func main() {

	var lessMonthIssues []*github.Issue
	var lessYearIssues []*github.Issue
	var plusYearIssues []*github.Issue

	t := time.Now()
	minus30Days := t.AddDate(0, -30, 0)
	minus1Year := t.AddDate(-1, 0, 0)

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range result.Items {
		if item.CreatedAt.After(minus30Days) {
			lessMonthIssues = append(lessMonthIssues, item)
		} else if item.CreatedAt.After(minus1Year) {
			lessYearIssues = append(lessYearIssues, item)
		} else {
			plusYearIssues = append(plusYearIssues, item)
		}
	}

	fmt.Println("Issues less than 30 days:")
	for _, item := range lessMonthIssues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("Issues less than 1 year:")
	for _, item := range lessYearIssues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("Issues over 1 year old:")
	for _, item := range plusYearIssues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
