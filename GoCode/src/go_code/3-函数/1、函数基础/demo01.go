package main

import (
	"fmt"
)

func main() {

	// 输入两个数，再输入一个运算符（+、-、*、/） 得到结果

	// 分析思路
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var result float64
	var operator byte = '*'
	
	switch operator {
		case '+':
			result = n1 + n2
		case '-':
			result = n1 - n2
		case '*':
			result = n1 * n2
		case '/':
			result = n1 / n2
		default:
			fmt.Println("操作符输入有误，无法计算")
	}
	fmt.Printf("%v %c %v = %v", n1, operator, n2, result)
}