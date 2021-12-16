package main

import (
	"fmt"
	"testing"
)

func Func_test(t *testing.T) {
	red := Redis_ctx{
		max_idle:    5,
		max_active:  5,
		idleTimeout: 5,
		host:        "127.0.0.1",
		timeout:     5,
	}
	red.Init()
	err := red.SSET("a", "b")
	if err != nil {
		fmt.Println(err)
	}
	value, err2 := red.SGET("a")
	if err2 != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}
