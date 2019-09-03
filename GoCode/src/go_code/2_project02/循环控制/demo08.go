package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// // 为了生成一个随机数，还需要为rand设置一个随机种子，否则生成的是一个固定的数（82）
	// // 调用time.Now().Unix()方法，返回一个从1970-01-01 00:00:00 到现在的秒数
	// rand.Seed(time.Now().Unix())
	// // 如何随机生成 1-100 整数
	// n := rand.Intn(100) + 1   // [1,100]随机整数
	// fmt.Println(n)

	// 编写一个无限循环，不停生成随机数，当生成了99时，退出循环
	count := 0
	for {
		rand.Seed(time.Now().UnixNano())  // 设置一个从1970-01-01 00:00:00 到现在的纳秒数作为随机种子
		n := rand.Intn(100) + 1
		fmt.Println("n=", n)
		count++
		if n == 99 {
			break   // 跳出for的死循环
		}
	}
	fmt.Printf("生成99，一共使用了%d次", count)
}