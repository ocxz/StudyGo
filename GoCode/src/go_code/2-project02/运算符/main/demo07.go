package main

import (
	"fmt"
)

func GetMax(num1 int , num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func GetMMax(num1 int, num2 int, num3 int) int {
	return GetMax(num3, GetMax(num1, num2))
}

func main() {

	// 演示 & 和 * 操作地址

	a := 100
	b := 200
	c := 300
	max := GetMax(a, b)
	fmt.Printf("%v 与 %v 的最大值是：%v \n", a, b, max)

	max = GetMMax(a, b, c)
	fmt.Printf("%v、%v 与%v 的最大值是：%v \n", a, b, c, max)
}