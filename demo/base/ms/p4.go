package main

import "fmt"

/**

  main和test调用&b和&test地址不同的原因是:新建变量，但是存储的指针是一样的
 */

func test(test map[string]string) {
	// test:0xc000078150
	fmt.Printf("test:%p\n", test)
	// &test:0xc000098028

	fmt.Printf("&test:%p\n", &test)
}

func main() {
	a := 10
	// a:%!p(int=10)
	fmt.Printf("a:%p\n", a)
	// &a:0xc000096010
	fmt.Printf("&a:%p\n", &a)
	//fmt.Println(&a)
	b := map[string]string{}
	// b:0xc000078150
	fmt.Printf("b:%p\n", b)
	// &b:0xc000098020
	fmt.Printf("&b:%p\n", &b)
	test(b)
	test(b)
}
