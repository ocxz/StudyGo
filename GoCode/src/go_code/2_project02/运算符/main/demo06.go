package main

import (
	"fmt"
)

func main() {

	// 演示 & 和 * 操作地址

	a := 100
	ptr := &a
	fmt.Printf("%v的地址是：%v \n", a, ptr)
	fmt.Printf("%v指向的值是：%v \n", ptr, *ptr)
}