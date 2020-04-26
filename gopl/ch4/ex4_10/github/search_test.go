package github

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"testing"
)

//测试url生成
//可以自己先把url沾到浏览器测试
func TestUrl(t *testing.T) {

	urls := []string{"repo:golang/go","is:open","json","decoder"}
	fmt.Printf("%s?+q=%s \n",IssuesURL,url.QueryEscape(strings.Join(urls," ")))
}
func TestSearchIssue(t *testing.T) {
	//term := []string{"gopl.io/ch4/issues"}

	term := []string{"repo:golang/go","is:open","json","decoder"}
	result,err := SearchIssue(term)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n",result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

