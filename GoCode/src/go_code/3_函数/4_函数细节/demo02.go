package main

import (
	"fmt"
)

// 在Go中，函数也是一种数据类型，
// 可以赋值给一个变量，该变量就是一个函数类型的变量，可通过变量对函数的调用
func GetSum(n1 int, n2 int) int {
	return n1 + n2
}

// 函数是一种数据类型，因此在Go中，函数可以作为形参且调用（类似委托）
func MyFunc(fn func(int, int) int, num1 int, num2 int) int{
	return fn(num1, num2)
}
func main() {

	// 在Go中，函数也是一种数据类型，
	// 可以赋值给一个变量，该变量就是一个函数类型的变量，可通过变量对函数的调用

	fn := GetSum
	fmt.Printf("fn的类型是%T，getSum的类型是%T \n", fn, GetSum)

	result := fn(10, 40)   // 等价于：GetSum(10, 20)
	fmt.Println("result=", result)

	result = MyFunc(fn, 10, 20)
	fmt.Println("result=", result)
}