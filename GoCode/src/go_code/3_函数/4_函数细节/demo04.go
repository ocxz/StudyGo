package main

import (
	"fmt"
)

// 编写一个sum，可以接受1到多个int类型参数，求和
func sum(n1 int, args... int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]   // 表示取出args切片的地i个元素，累加到sum中
	}
	return sum
}
func main() {

	// 测试函数的可变参数
	// 演示可变参数的使用
	result := sum(4)
	fmt.Println("result=", result)

	result = sum(4, 5, 6, 7)
	fmt.Println("result=", result)
}