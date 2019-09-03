package main

import (
	"fmt"
)

func main() {

	// 演示一下指定标签形式来使用break
	// level1:
	for i:= 0; i < 4; i++ {
		level2:    // 设置一个标签，名字随意
		for j:= 0; j < 10; j++ {
			if j == 2 {
				// break   // 默认跳出最近的for循环
				// break level1
				break level2
			}
			fmt.Println("j=", j)
		}
	}
}