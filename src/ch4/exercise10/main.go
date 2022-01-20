// Issues prints a table of Github issues matching the search terms.
// go run main.go repo:golang/go is:open json decoder
package main

import (
	"fmt"
	"go-programming-language/src/ch4/github"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	sort.Slice(result.Items, func(i, j int) bool {
		return result.Items[i].CreatedAt.After(result.Items[j].CreatedAt)
	})

	now := time.Now()

	sixMonthsBefore := now.AddDate(0, -6, 0)
	lessThanSixMonthsOld := sort.Search(len(result.Items), func(i int) bool {
		return result.Items[i].CreatedAt.Before(sixMonthsBefore)
	})
	fmt.Println("*** issues < 6 months ***")
	PrintIssues(result.Items[:lessThanSixMonthsOld])

	oneYearBefore := now.AddDate(-1, 0, 0)
	oneYearOld := sort.Search(len(result.Items), func(i int) bool {
		return result.Items[i].CreatedAt.Before(oneYearBefore)
	})
	fmt.Println("*** issues 6 months - 1 year ***")
	PrintIssues(result.Items[lessThanSixMonthsOld:oneYearOld])

	fmt.Println("*** issues > 1 year ***")
	PrintIssues(result.Items[oneYearOld:])

	fmt.Println("*** all issues ***")
	PrintIssues(result.Items)
}

func PrintIssues(items []*github.Issue) {
	for _, item := range items {
		PrintIssueDetails(item)
	}
}

func PrintIssueDetails(item *github.Issue) {
	fmt.Printf("#%-5d %s %9.9s %.55s\n",
		item.Number, item.CreatedAt, item.User.Login, item.Title)
}
