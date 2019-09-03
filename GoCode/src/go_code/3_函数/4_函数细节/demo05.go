package main

import (
	"fmt"
)

// 传入两个整数变量地址，交换两个变量的值
func swap(n1, n2 *int) {
	temp := *n1  // 定义一个temp变量，存储n1指针指向的值
	*n1 = *n2
	*n2 = temp
}
func main() {

	n1, n2 := 10, 20
	fmt.Printf("交换前：n1=%v，n2=%v \n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("交换后：n1=%v，n2=%v \n", n1, n2)
}