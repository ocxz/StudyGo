package main

import (
	"fmt"
)

func main() {

	// continue 小练习
	// continue实现 打印1-100 之内的奇数，要求使用for+continue
	for i := 1; i <= 100; i++ {
		if i % 2 == 0 {
			continue
		} else {
			fmt.Println("奇书是：", i)
		}
	}
}