package github

import (
	"fmt"
	"log"
	"testing"
)

func TestSearchByDate(t *testing.T)  {
	term := []string{"repo:golang/go","is:open","json","decoder"}
	result,err := SearchIssue(term)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(result)
	searchByDate(result.Items)
}
