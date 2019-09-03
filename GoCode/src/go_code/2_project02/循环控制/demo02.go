package main

import (
	"fmt"
)

func main() {

	// for 循环小练习
	// 1、打印1-100之间所有9的倍数个数、及总和
	count, sum := 0, 0
	for i := 9; i <= 100; i+=9 {
		// if i % 9 == 0 {
		// 	count += 1
		// 	sum += i
		// }

		count += 1
		sum += i
	}
	fmt.Printf("1-100内所有9倍数的个数是：%d个, 总和是：%d", count, sum)
	fmt.Println()

	// 2、 完成 0 + 6 = 6   1 + 5 = 6 ……  6 + 0 = 6 的输出
	var num int
	fmt.Print("请输入一个正整数：")
	fmt.Scanln(&num)

	for i := 0; i <= num; i++ {
		fmt.Printf("%d\t+\t%d\t=\t%d\n", i, num-i, num)
	}
}