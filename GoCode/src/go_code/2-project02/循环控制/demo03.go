package main

import (
	"fmt"
)

func main() {

	// 使用for循环 模拟while循环
	i := 1   // 初始化循环变量
	for {
		if i > 10 {   // 循环条件
			break  // 跳出该for循环
		}

		fmt.Println("hello world ", i)  // 循环体
		i++    // 循环迭代
	}
	fmt.Println()

	// 使用for循环 模拟do...while循环
	i = 1
	for {
		fmt.Println("hello world", i)
		i++
		if i > 10 {
			break   // 就是跳出for循环
		}
	}
}