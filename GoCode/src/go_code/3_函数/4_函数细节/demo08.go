package main

import (
	"fmt"
)

// 定义全局变量，全局匿名函数
var (
	// fun1 就是一个全局匿名函数
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {

	// 匿名函数 一种 在定义时直接调用，只能使用一次
	// 案例演示 使用匿名函数求两个数的和
	result1 := func (n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("result1=",result1)

	// 第二种 将匿名函数赋值给函数变量 通过函数变量实现多次调用
	// 案例演示  使用匿名函数 函数变量 实现两个数的差
	sub := func(n1 int, n2 int) int {
		return n1 - n2
	}
	result2 := sub(10, 20)
	fmt.Println("result2=", result2)

	// 全局匿名函数的使用
	result3 := Fun1(4, 9)
	fmt.Println("result3=", result3)
}