package main

import (
	"fmt"
	"sync"
)

//TODO 多运行几次  有错误 但是不理解
var set = make(map[int]bool)
var waits sync.WaitGroup
var count int
func printOnce(num int){
	if _,exit := set[num];!exit{
		count++
		fmt.Println(num)
	}
	set[num] = true
	waits.Done()
}
func main(){
	waits.Add(100)
	for i :=0;i<100;i++ {
		
		go printOnce(100)
	}
	waits.Wait()
	fmt.Println(count)
}