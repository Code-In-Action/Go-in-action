package search

import (
	"log"
	"sync"
)

// A map of registered matchers for searching.
var matchers = make(map[string]Matcher)

// searchTerm 参数 string类型
func Run(searchTerm string) {
	// feeds, err 是RetrieveFeeds()函数的返回值， err是错误值
	// 简化声明变量，和用 var的没啥区别
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	results := make(chan *Result)
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds));

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}
		go func (matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feedd)

	}
	go func() {
		waitGroup.Wait()
		close(results)
	}()
	Display(results)
}