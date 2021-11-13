package main

import "fmt"

var ch = make(chan string)

func main() {
	fmt.Println(len(ch))
	push("a")
	fmt.Println(pop())
	fmt.Println(pop())
	fmt.Println(pop())
	push("a")
	push("a")
	push("a")
}

func push(s string) {
	ch <- s
}

func pop() string {
	return <-ch
}
