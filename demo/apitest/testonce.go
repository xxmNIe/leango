package main

import (
	"fmt"
	"sync"
)

var once sync.Once
func main() {
	onceBody := func(){
		fmt.Println("only once")
	}

	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
