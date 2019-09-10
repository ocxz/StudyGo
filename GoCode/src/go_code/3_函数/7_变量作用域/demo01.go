package main

import (
	"fmt"
)

var gender string = "男"
// 函数
func test() {

	// 作用域仅限于本函数内部使用
	age := 10
	Name := "tom"
	fmt.Printf("%v是一个%v生，今年%v岁了 \n", Name, gender, age)
}

var name = "tom"  // 全局变量
func test01() {
	fmt.Println(name)
}
func test02() {
	// 就近原则
	name := "jack"
	fmt.Println(name)
}

func main() {
	// test(gender)
	test()
	
	fmt.Println(name)   // tom
	test01()    // tom
	test02()    // jack
	test01()    // tom
}