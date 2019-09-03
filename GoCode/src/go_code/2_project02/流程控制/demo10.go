package main

import (
	"fmt"
)

func main() {

	// 1、判断一个整数是否为水仙花数，要求各个数字立方之和等于本身 如：153 = 1^3 + 5^3 + 3^3
	// 解决关键：求出该三位数的各个数字
	for {
		var num int
		fmt.Print("请输入一个三位正整数：")
		fmt.Scanln(&num)
		if num < 100 || num > 999 {
			fmt.Println("输入有误，无法判断")
		} else {
			n1 := num / 100     // 取百位
			n2 := (num - 100 * n1) / 10   // 取十位
			n3 := num - n1 *100 - n2 * 10   // 取个位
	
			if num == n1*n1*n1 + n2*n2*n2 + n3*n3*n3 {
				fmt.Printf("%d = %d^3 + %d^3 + %d^3 是水仙花数 \n", num, n1, n2, n3)
			}else {
				fmt.Printf("%d不是水仙花数 \n", num)
			}
		}
	}
}