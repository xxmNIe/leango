package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup
func main()  {
	var urls =[]string{
		"http://www.baidu.com",
		"http://bilibili.com",
		"http://jandan.net/",
	}

	for _,url := range urls {
		wg.Add(1)
		time.Sleep(1*time.Second)
		go func(url string) {
			defer wg.Done()
			time.Sleep(1*time.Second)
			http.Get(url)
		}(url)
	}
	wg.Wait()
	fmt.Println("ojbk")
}
