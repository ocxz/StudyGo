package main

import (
	"fmt"
)

func main() {

	// continue小练习
	// 从键盘输入数，判断正数、负数个数，输入为0时结束程序
	// 使用for循环、break、continue完成

	posCount, negCount := 0, 0   // 分别记录正数个数、负数个数

	forLabel:
	for {
		var num float64
		fmt.Print("请输入一个数：")
		fmt.Scanln(&num)

		switch {
		case num > 0 :
			posCount++
		case num < 0 :
			negCount++
		default:
			break forLabel
		}
	}
	fmt.Printf("一共输入了%d个正数，%d个负数 \n", posCount, negCount)
}