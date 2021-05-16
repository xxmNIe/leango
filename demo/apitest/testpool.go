package main

import (
	"fmt"
	"sync"
)

func main() {
	p := &sync.Pool{
		New: func() interface{}{return 1 },
		}
	p.Put(29)
	p.Put(22)
	p.Put(28)
	p.Put(25)
	p.Put(4)
	p.Put(3)

	fmt.Println(p.Get())
	fmt.Println(p.Get())
	fmt.Println(p.Get())
	fmt.Println(p.Get())


}
