package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
	基本类型和数组都是值类型
	切片引用类型
	 */
	fmt.Println("------------slice-------------")
	a := []int{1,2,3}
	b :=a
	b[1]=44
	fmt.Println(a)
	fmt.Println("------------map-------------")
	m1:= map[int]int{1:1,2:2,3:3}
	m2 := m1
	m2[1] = 11111
	fmt.Println(m1)


	fmt.Println("                             ")
	fmt.Println("                             ")
	fmt.Println("------------值类型------------")
	fmt.Println("------------array-------------")
	t:= [3]int{11,22,44}
	m := t
	m[1]=9090
	fmt.Println(t)


}

/*
make new 切片的区别
 */
func test1(){
	a := new([]int)
	b := make([]int,1)
	fmt.Println(a,b)
	fmt.Println(reflect.TypeOf(a),reflect.ValueOf(a))
	fmt.Println(reflect.TypeOf(b),reflect.ValueOf(b))
	//*a 取到a的内容 这时没有为底层初始化数组,所以只是指针的零值
	fmt.Printf("pa:%p  pb:%p\n",*a,&b)

	a = &[]int{1,2,3}
	fmt.Println(reflect.TypeOf(a),reflect.ValueOf(a))
	fmt.Println(reflect.TypeOf(b),reflect.ValueOf(b))
	fmt.Printf("pa:%p  pb:%p\n",*a,&b)
	/*
		&[] [0]
		*[]int &[]
		[]int [0]
		pa:0x0  pb:0xc00011c030
		*[]int &[1 2 3]
		[]int [0]
		pa:0xc000134000  pb:0xc00011c030

	*/
}