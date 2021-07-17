package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("前")
	}()
	defer func() {
		if err :=recover();err !=nil {
			fmt.Println(err)
		}
	}()
	defer func() {
		fmt.Println("中")
	}()
	defer func() {
		fmt.Println("后")
	}()

	panic("paaaaa")
}
