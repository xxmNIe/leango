package main

import "fmt"

func main() {
	//ch := make(chan int)
	var  ch chan int
	//for i := 0; i < 1; i++ {
	//	go func(idx int) {
	//		ch <- idx
	//	}(i)
	//
	//}
	go func() {
		select {
		case s := <-ch:
			fmt.Println(s)

		}
	}()
	ch <- 1
	//close(ch)


}