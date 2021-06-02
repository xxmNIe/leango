package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string)
	go func() {


		ch1 <- "hello world"
		ch1 <- "hello China"
	}()

	fmt.Println(<-ch1)
	select {

	}

}