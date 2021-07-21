package main

import (
	"fmt"
	"sync"
)

func main() {
	smp :=sync.Map{}
	mp := map[int]int{}
	mp[1] = 2
	go func() {
		for  {
			//mp[1] =3
			smp.Store(1,3)
		}

	}()
	go func() {
		for  {
			//mp[2] =4
			smp.Store(1,4)
		}

	}()
	go func() {
		for {
			//m1 := mp[1]
			//m2 := mp[1]
			m1,_:=smp.Load(1)
			m2,_:= smp.Load(1)
			if m1 !=m2 {
				fmt.Println("no","   ",m1,m2)
			}else {
				fmt.Println("yes    ",m1,m2)
			}
		}
	}()
	//go func() {
	//	for {
	//		m1 := mp[2]
	//		m2 := mp[2]
	//		if m1 !=m2 {
	//			fmt.Println("no","   ",m1,m2)
	//		}else {
	//			fmt.Println("yes    ",m1,m2)
	//		}
	//	}
	//}()
	select {

	}
}
