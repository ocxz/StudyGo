package main

import (
	"fmt"
)

func main(){

	// 两个变量、进行交换、不用中间变量
	// var a, b = 10, 20
	// a, b := 10, 20
	var a int = 10

	// 将a与b的和保持  a = a + b
	// b = 两者之和减去原b 等到结果为原a   b = a - b
	// a = 两者之和将去原a即b  得到结果为b  a = a - b
	a = a + b
	b = a - b
	a = a -b 
	fmt.Printf("a=%v，b=%v \n", a, b)

	

	// 将两者之差保存起来  a = a - b
	// b = 原b + 两者之差  等到结果为原a  b = b + a
	// a = 原a - 两者之差  得到结果为原b  a = b - a
	a = a - b
	b = b + a  // b
	a = b - a
	fmt.Printf("a=%v，b=%v \n", a, b)

}