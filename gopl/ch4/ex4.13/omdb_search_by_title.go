package ex4_13

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

const apikey = "837a1b8b"
const apiUrl = "http://www.omdbapi.com/"

type moveInfo struct {
	Title  string
	Year   string
	Poster string
}

//福尔摩斯
// OMDb API: http://www.omdbapi.com/?t=Holmes&apikey=837a1b8b
func getPoster(title string) {
	resp, err := http.Get(fmt.Sprintf("%s?t=%s&apikey=%s", apiUrl, url.QueryEscape(title), apikey))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	binfo, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	minfo := moveInfo{}
	err = json.Unmarshal(binfo, &minfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	poster := minfo.Poster
	if poster != "" {
		downloadPoster(poster)
	}
}
func downloadPoster(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	bcontent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	pos := strings.LastIndex(url, "/")
	if pos == -1 {
		fmt.Println("failed to parse the title of the poster")
		return
	}
	f, err := os.Create(url[pos+1:])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	_, err = f.Write(bcontent)
	if err != nil {
		fmt.Println(err)
	}
}
func searchByTitle(titles ...string) {
	var wg sync.WaitGroup
	for _, title := range titles {
		wg.Add(1)
		go func() {
			getPoster(title)
			wg.Done()
		}()
	}
	wg.Wait()
}

