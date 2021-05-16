package main

import "fmt"

type Handler interface {
	Do(k,v interface{})
}

type HandleFunc func(k,v interface{})

func (hf HandleFunc) DO(k,v interface{}) {
	hf(k,v)
}

func Each(mp map[interface{}]interface{},h Handler) {
	if mp !=nil && len(mp) >0 {
		for k,v := range mp {
			h.Do(k,v)
		}
	}
}

func EachFunc(mp map[interface{}]interface{}, handlerFunc HandleFunc) {
	if mp != nil && len(mp) > 0 {
		for k,v := range mp {
			handlerFunc(k,v)
		}
	}
}

func selfInfo(k,v interface{}) {
	fmt.Printf("my name is %s, i am %d years old",k,v)
	fmt.Println()
}

func main() {
	mp := map[interface{}]interface{}{
		"gaoziwen":26,
		"zhangsan":27,
		"lisi":28,
	}
	EachFunc(mp,selfInfo)

}