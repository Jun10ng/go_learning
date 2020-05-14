package search

import (
	"log"
	"sync"
)

//注册用于搜索的匹配器map
var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个无缓冲的通道。接手匹配后的结果
	results := make(chan *Result)

	// waitGroup 处理所有数据源
	var wg sync.WaitGroup

	// 设置任务数
	wg.Add(len(feeds))

	//为每个数据源启动一个gorountine
	for _, feed := range feeds {
		// 获取一个匹配器用于查找
		matcher, ok := matchers[feed.Type]
		if !ok {
			matcher = matchers["default"]
		}

		//启动goroutine执行搜索
		go func(matcher Matcher, feed *Feed) {
			// defer wg.Done()
			Match(matcher, feed, searchTerm, results)
			wg.Done()
		}(matcher, feed)
	}
	// 启动一个协程用于监控是否所有工作都做完了
	go func() {
		wg.Wait()
		close(results)
	}()
	// 启动函数，显示返回的结果，并且在最后一个结果显示完后返回
	Display(results)
}

// Register 调用时会注册一个匹配器，提供给后面的程序使用
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatal(feedType, "Matcher already registered")
	}
	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
