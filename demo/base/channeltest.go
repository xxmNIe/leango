package main

import (
	"fmt"
	"time"
)



func deadlock(){
	//pipe := make(chan int,3)
	pipe := make(chan int,3)
	go func(){<-pipe}()
	pipe <- 1
	fmt.Println("1")
	pipe <- 2
	fmt.Println("2")
	pipe <- 3
	fmt.Println("3")
}


func closed(){
	pipe:=make(chan int)

	go func() {
		fmt.Println("    send 1")
		pipe <- 1
		time.Sleep(1*time.Second)
		fmt.Println("    send 1")
		pipe <-1
		time.Sleep(1*time.Second)
		fmt.Println("close")
		close(pipe)

		time.Sleep(1*time.Second)
		fmt.Println("    send 2")
		pipe <-2

	}()

	for  {
		ch,ok := <-pipe
		fmt.Println(ok,"  ",ch)
		time.Sleep(1*time.Second)
	}

}
func main() {

	closed()
}