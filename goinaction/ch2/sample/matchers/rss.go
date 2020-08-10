package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"go_learning/goinaction/ch1/sample/search"
	"log"
	"net/http"
	"regexp"
)

type (
	// item 根据item字段的标签，将定义的字段与rss文档的字段关联起来
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}
	//image根据image字段的标签，将定义的字段与rss文档的字段连接起来
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// channel 根据 channel 字段的标签，将定义的字段与rss文档的字段关联起来
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuilder"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument 定义了与rss 文档相关联的字段
	rssDocument struct {
		XMLName xml.Name `xml:"rss" `
		Channel channel  `xml:"channel"`
	}
)


// rssMatcher 实现了Matcher接口
type rssMatcher struct{}

// init 匹配注册到程序里
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// Search 在文档中查找特定的搜索项
func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result

	log.Printf("Search Feed Type [%s] Site[%s] For Uri[%s]\n", feed.Type, feed.Name, feed.URI)

	// 获取要搜索的数据
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}
	for _, channelItem := range document.Channel.Item {
		// 检查标题部分是否包含搜索项
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}
		// 如果找到匹配的项，将其作为结果保存
		if matched {
			results = append(results, &search.Result{
				Field:   "Title",
				Content: channelItem.Title,
			})
		}

		// 检查描述部分是否包含搜索项
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		// If we found a match save the result.
		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}
	return results, nil
}
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed URI provided")
	}
	// 从网络获得rss数据源文档
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	// 一旦从函数返回，关闭返回的响应连接
	defer resp.Body.Close()
	// 检查状态码是不是200，这样就能知道是不是收到正确的响应了

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("URI[%s]HTTP Response Error %d\n", feed.URI, resp.StatusCode)
	}
	// 将 rss 数据源文档解码到我们定义的结构类型里不需要检查错误，调用者会做
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}