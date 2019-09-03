package main

import (
	"fmt"
)

// prt 是指针类型，即传入的是一个int数的地址
func test (prt *int) {

	*prt = *prt + 20 // 通过指针操作变量，来修改函数外的变量值
}
func main() {

	// 可传入变量地址给函数，使函数内可通过指针操作变量来修改函数外的变量值
	n := 10
	test(&n)
	fmt.Println("n=", n)
}