package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {

	var  buffSize int
	var  consumers int
 	var  producter int
	var  du time.Duration

	flag.IntVar(&buffSize,"b",10,"缓冲区大小")
	flag.IntVar(&consumers,"c",5,"消费者数量")
	flag.IntVar(&producter,"p",5,"生产者数量")
	flag.DurationVar(&du,"t",5,"运行时间")
	buff := make(chan interface{},buffSize)
	var wg sync.WaitGroup
	final := time.After(time.Second*du)
	wg.Add(consumers)
	for i:=0;i<consumers;i++ {
		cNum :=i
		go func(q <-chan interface{},num int,waitgroup *sync.WaitGroup) {
			for _ =range q {
				fmt.Println("consumer:[", num, "] ", "consume a product")
			}
			waitgroup.Done()
		}(buff,cNum,&wg)
	}

	for i:=0;i<producter;i++ {
		cNum :=i
		go func(q  chan<- interface{},num int) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("prouduct ", cNum,"Done ")
				}
			}()
			for {
				q<- struct {}{}
				fmt.Println("producter:[",num,"] " ,"product a product")
			}

		}(buff,cNum)
	}
	select {
	case  <-final:
		close(buff)
	}
	wg.Wait()
}