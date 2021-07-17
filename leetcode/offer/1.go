package main

import (
	"container/list"
	"fmt"
)

//type CQueue struct {
//	s1 []int
//	s2 []int
//}
//
//func Constructor() CQueue {
//	return CQueue{
//		s1 : make([]int,0),
//		s2 : make([]int,0),
//	}
//}
//
//
//func (this *CQueue) AppendTail(value int)  {
//	this.s1 =  append(this.s1, value)
//
//}
//
//
//func (this *CQueue) DeleteHead() int {
//	if len(this.s1)==0 {
//		return -1
//	}
//
//	i:=0
//	for i<len(this.s1) {
//		if i<len(this.s2) {
//			for i<len(this.s2){
//				this.s2[i] = this.s1[i]
//				i++
//			}
//		}else {
//			this.s2 = append(this.s2,this.s1[i])
//		}
//
//	}
//
//	res:= this.s2[len(this.s2)-1]
//	i=0
//	for i<len(this.s2)-1{
//		this.s1[i] = this.s2[i]
//	}
//	return res
//}

type CQueue struct {
	s1,s2 *list.List
}


func Constructor() CQueue {
	return CQueue{
		s1 : list.New(),
		s2 : list.New(),
	}
}


func (this *CQueue) AppendTail(value int)  {
	this.s1.PushBack(value)
}


func (this *CQueue) DeleteHead() int {
	if this.s2.Len()  == 0 {
		for this.s1.Len() > 0 {
			this.s2.PushBack(this.s1.Remove(this.s1.Back()))
		}
	}
	if this.s2.Len()!=0 {
		e :=this.s2.Back()
		this.s2.Remove(e)
		return e.Value.(int)
	}
	return -1
}




func main() {
	 obj := Constructor()
	 obj.AppendTail(23)
	 param_2 := obj.DeleteHead()
//	param_2 = obj.DeleteHead()
	 fmt.Println(param_2)
}