package main

import (
	"fmt"
	"reflect"
)

type person interface {
	AddFace(face Face)
}

type Face struct {
	Size string
	len  int
	in   inner
}
type inner struct {
	a int
}

type Man struct {
	face Face
	list []int
}

func (m *Man) AddFace(face Face) {
	m.face = face
}
func display(man Man) {
	man.list = make([]int, 4)
	fmt.Println(man)
}
func main1() {
	//字面量书初始化结构体 返回引用 内嵌结构体基本类型为初始零值，
	//给结构体赋值后无论成员是结构体还是其他类型,只分基本类型和指针类型   基本复制,外部引用修改不会修改结构体内的值 指针会
	//传入的内嵌结构体在外部的引用改变该内嵌结构体的基本类型参数时，结构体内的内嵌结构体不收受影响
	//结构体作为参数传入,修改内部的slice map(指针类型)也会引起调用者的变化  所以需要改变基本类型时候传引用,否则传值,还要注意浅拷贝带来的问题
	m := Man{}
	inner := inner{2}
	face := Face{"b", 2, inner}
	m.AddFace(face)

	tmp := make([]int, 2)
	m.list = tmp
	m.list[0] = 3
	fmt.Println(reflect.TypeOf(m))
	fmt.Println(reflect.TypeOf(face))
	inner.a = 4
	//浅拷贝  slice map会共享，
	x := m
	display(x)
	m.list[1] = 2
	m.face.len = 555
	tmp[0] = 1
	display(x)
	display(m)

}
