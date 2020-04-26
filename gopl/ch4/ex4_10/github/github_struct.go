

package github

import (
	"time"
)

const IssuesURL  = "https://api.github.com/search/issues"

type Issue struct {
	Number 		int
	HTMLURL 	string `json:"html_url"`
	Title 		string
	State 		string
	User        *User
	CreatedAt	time.Time `json:"created_at"`
	Body 		string
}
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items			[]*Issue
}

type User struct {
	Login 	string
	HTMLURL	string `json:"html_url"`
}