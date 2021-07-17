package main

import "fmt"

func main() {
	switch "a" {
	case "a":
		fmt.Println("匹配a")
		fallthrough
	case "b":
		fmt.Println("a成功了，也执行b分支")
		fallthrough
	case "c":
		fmt.Println("a成功了，c分支会执行吗？")
		fallthrough
	default:
		fmt.Println("默认执行")
	}

}
