package main

import (
	"flag"
	"fmt"
)

func main() {
	rPtr := flag.Bool("r",false,"kernel release")
	aPtr := flag.Bool("a",false,"all info")
	mPtr := flag.Bool("m",false,"machine")
	sPtr := flag.Bool("s",false,"kernel name")

	flag.Parse()

	if(*mPtr){
		fmt.Println("x86_64")
	}else if(*aPtr){
		fmt.Println("Linux localhost 4.4.254-36964-g86d266acc51a #1 SMP PREEMPT Mon May 10 21:47:52 PDT 2021 x86_64 x86_64 x86_64 GNU/Linux")
	}else if(*rPtr){
		fmt.Println("4.4.254-36964-g86d266acc51a")
	}else if(*sPtr){
		fmt.Println("linux")
	}else{
		fmt.Println("linux")
	}
}
