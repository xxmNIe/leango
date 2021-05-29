package main

import (
	"fmt"
	"reflect"
)



func main() {
	//type T1 struct {
	//	name string
	//}
	//type T2 struct {
	//	name string
	//}
	//t1 := T1{"abc"}
	//t11 := T1{"abc"}
	//t2 := T2{"abc"}
	//t22 := T2{"abc"}
	////Invalid operation: t1 == t2 (mismatched types T1 and T2)
	////fmt.Printf("%v",t1 == t2)
	//fmt.Println(t1 == t11)
	////t11.name = "abbbb"
	//fmt.Println(t1 == t11)
	//// falsexz`
	//fmt.Printf("t1 adrr :%p,  t11 addr :%p\n",&t1,&t11 )
	////false
	//fmt.Println(t2 == t22)

/*
	不同类型的不可比较
	关于类型相同
 */
	//type Int int
	//var v11 int = 1
	//var v12 int = 1
	//v13 :=&v12
	//var v21 Int = 1
	//var v22 Int = 1
	//
	//fmt.Printf("v11 == v22: %v\n",v11 == v12)
	////都不相同
	//fmt.Printf("addr v11:%p   v12:%p v13:%p\n",&v11 ,&v12,v13)
	////Invalid operation: v11==v21 (mismatched types int and Int)
	////fmt.Printf("v11 == v21: %v",v11==v21)
	//fmt.Printf("v21==v22: %v\n",v21==v22)
	//fmt.Printf("addr v21:%p   v22:%p\n",&v21 ,&v22)

	fmt.Println("-----------------------------------------    ----------------------------")
	/*
		仅当struct的所有字段都可比较时，struct才可比较
	 */
	//
	//type s1 struct {
	//	name string
	//}
	//
	//type s2 struct {
	//	name string
	//	hob map[string]interface{}
	//}
	//
	//v11 := s1 { "foo" }
	//v12 := s1 { "foo" }
	//v21 := s2 { "foo", make(map[string]interface{}) }
	//v22 := s2 { "foo", make(map[string]interface{}) }
	//fmt.Printf("v11 == v12 is %v\n", v11 == v12)
	//Invalid operation: v21 == v22 (the operator == is not defined on s2)
	//fmt.Printf("v21 == v22 is %v\n", v21 == v22)


	//匿名类型
	type T1 struct {
		name string
	}
	type T2 struct {
		name string
	}

	t1 := T1{"abc"}
	t2 := T2{"abc"}
	t3 := struct {
		name string
	}{"abc"}
	t4 := struct {
		name string
	}{"abc"}

	fmt.Printf("t1 type:%v  t1 value:%v\n",reflect.TypeOf(t1),reflect.ValueOf(t1))
	fmt.Printf("t2 type:%v  t2 value:%v\n",reflect.TypeOf(t2),reflect.ValueOf(t2))
	fmt.Printf("t3 type:%v  t3 value:%v\n",reflect.TypeOf(t3),reflect.ValueOf(t3))
	fmt.Printf("t4 type:%v  t4 value:%v\n",reflect.TypeOf(t4),reflect.ValueOf(t4))

	//all  ture
	fmt.Println(t1 == t3)
	fmt.Println(t2 == t3)
	fmt.Println(t3 == t4)
}