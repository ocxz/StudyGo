package main

import (
	"fmt"
)

// 演示golang指针类型的问题
func main() {

	// 基本数据类在内存中的布局
	var i int = 10
	// i 的地址取出  可通过&i 取出i所在内存区域的地址
	fmt.Println("i的地址=", &i)

	// 声明指针变量，且给指针变量赋值
	var ptr *int = &i
	fmt.Println("ptr指针变量的值=", ptr)
	fmt.Println("ptr的地址=", &ptr)

	// 可以使用* 取出指针所指向的值
	fmt.Println("ptr指针所指向的值=", *ptr)
}
