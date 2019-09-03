package main

import (
	"fmt"
	"go_code/3-函数/1_函数基础/utils"
)

func main() {

	// 输入两个数，再输入一个运算符（+、-、*、/） 得到结果

	// 分析思路
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte = '*'

	result := utils.Calculate(n1, n2, operator)
	fmt.Printf("%v %c %v = %v", n1, operator, n2, result)
}