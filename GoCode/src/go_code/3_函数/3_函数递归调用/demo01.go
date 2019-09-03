package main

import (
	"fmt"
)

func test(n int) {
	if n > 2 {
		n--  // 递归必须向退出递归条件逼近，否则就是无限循环调用，很可怕
		test(n)
	}
	fmt.Println("n=", n)
}

func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println("n=", n)
	}
}

func main() {

	test2(4)
	// fmt.Println("hello world")
}