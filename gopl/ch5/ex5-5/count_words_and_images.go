package ex5_5

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"strings"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	u := make([]*html.Node, 0) // unvisited
	u = append(u, n)
	for len(u) > 0 {
		n = u[len(u)-1]
		u = u[:len(u)-1]
		switch n.Type {
		case html.TextNode:
			words += wordCount(n.Data)
		case html.ElementNode:
			if n.Data == "img" {
				images++
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			u = append(u, c)
		}
	}
	return
}

func wordCount(s string) int {
	n := 0
	scan := bufio.NewScanner(strings.NewReader(s))
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		n++
	}
	return n
}

func cWaI(url string) {

	words, images, err := CountWordsAndImages(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("words: %d\nimages: %d\n", words, images)
}