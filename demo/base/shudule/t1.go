package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var x int
	threads := runtime.GOMAXPROCS(0)
	for i := 0; i < threads; i++ {
		go func() {
			// for {
			// 	x++
			// }
			x++	
		}()
	}
	x= x+1
	time.Sleep(time.Second)
	fmt.Println("x =", x)
}
