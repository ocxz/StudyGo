package main

import (
	"fmt"
)

func main() {

	// continue 案例分析

	level1:
	for i:= 0; i < 4; i++ {
		// level2:    // 设置一个标签，名字随意
		for j:= 0; j < 10; j++ {
			if j == 2 {
				// continue   // 默认跳出最近的for循环
				continue level1
				// continue level2
			}
			fmt.Println("j=", j)
		}
	}
}