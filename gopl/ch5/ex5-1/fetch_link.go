package ex5_1

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)



func fetchLink(url string){
	// fetch
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	//b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	//fetch link
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	visit(nil,doc)
	//for _, link := range visit(nil, doc) {
	//	fmt.Println(link)
	//}
}

//!-main

//递归调用
func visit(links []string, n *html.Node){
	//递归出口
	if n==nil{
		return
	}
	//真正流程
	visit(links,n.FirstChild)
	visit(links,n.NextSibling)
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a.Val)
				links = append(links, a.Val)
			}
		}
	}
	//for c := n.FirstChild; c != nil; c = c.NextSibling {
	//	links = visit(links, c)
	//}
}

//!-visit

/*
//!+html
package html
type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}
type NodeType int32
const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)
type Attribute struct {
	Key, Val string
}
func Parse(r io.Reader) (*Node, error)
//!-html
*/