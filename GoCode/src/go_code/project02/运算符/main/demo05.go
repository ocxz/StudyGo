package main

import (
	"fmt"
)

func main() {

	a := 2   // 二进制为：10
	b := 3   // 二进制为：11

	c := a & b  // 即 10 & 11 => 10 结果为2
	fmt.Printf("%v & %v = %v \n", a, b, c)

	d := a | b  // 即 10 & 11 => 11 结果为3
	fmt.Printf("%v & %v = %v \n", a, b, d)

	e := a ^ b  // 即 10^ 11 ==> 01 结果为：1
	fmt.Printf("%v & %v = %v \n", a, b, e)

	// << 向左移位多少 如：二进制10左移位2 即-> 1000 结果为2^3 = 811
	// 格式：a << 2 表示 a 左移2位
	// 左移多少位，就乘以2的多少次方
	f := a << 2
	fmt.Printf("%v << 2 = %v \n", a, f)


	// 右移多少位，就除以2的多少次方
	m := f >> 1
	fmt.Printf("%v >> 1 = %v \n", f, m)

	fmt.Println("-2^3=", -2^3)
}