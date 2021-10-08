package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchItem string) {
	feeds, err := RetrieveFeed()
	if err != nil {
		log.Fatal(err)
	}
	//创建一个无缓冲的通道,用于接受匹配后的结果
	results := make(chan *Result)

	//WaitGroup类似Java的CountDownLatch
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	//为每个feed产生一个goroutine去查找
	for _, feed := range feeds {
		matcher, exist := matchers[feed.Type]
		if !exist {
			matcher = matchers["default"]
		}

		//启动一个goroutine进行搜索
		go func(matcher Matcher, feed *Feed) {
			//TODO 调用匹配函数
			Match(matcher, feed, searchItem, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	go func() {
		//等待所有goroutine执行完毕
		waitGroup.Wait()
		//关闭通道
		close(results)
	}()

	//TODO 调用展示结果的函数
	Display(results)
}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}
	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
