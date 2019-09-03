package main

import (
	"fmt"
)

// 演示golang的基本数据类型转换问题
func main() {

	var i int = 100

	// 将i转成float32类型
	var f float32 = float32(i)

	// fmt.Printf("i=%T, f=%T", i, f)
	fmt.Println("i=", i, "f=", f)


	var j int64 = 100
	var k int8 = int8(j)
	fmt.Println(k)

	var o int64 = 1000000000
	var m int8 = int8(o)
	fmt.Println("m=", m)

}
