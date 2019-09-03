package main

import (
	"fmt"
)

// 演示golang的基本数据类型转换问题, 练习
func main() {

	var n1 int32 = 12
	var n2 int64
	var n3 int8
	n2 = int64(n1) + 20
	n3 = int8(n1) + int8(n2)

	fmt.Println("n2=", n2, "n3=", n3)

	// var n4 int8 = 129 - int8(n1)
	var n4 int8 = 129 - 18
	fmt.Println("n4=", n4)
}
