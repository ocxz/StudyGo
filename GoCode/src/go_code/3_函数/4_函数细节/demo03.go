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

// 在举一个案例
type myFun func(int, int) int    // 这时 myFunc 就是 func(int, int) int类型了
func MyFunc2(fn myFun, num1 int, num2 int) int{
	return fn(num1, num2)
}

// 支持对函数返回值命名
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sub = n1 - n2
	sum = n1 + n2
	return
}
func main() {

	// 看案例
	type myInt int  // 给int取了别名  在go中 myInt和int虽然都是int类型，但go认为是两个类型
	var num myInt
	var num2 int
	num = 40
	num2 = int(num)
	fmt.Println("num=", num)
	fmt.Println("num2=", num2)

	// // 在举一个案例
	// type myFun func(int, int) int    // 这时 myFunc 就是 func(int, int) int类型了
	result := MyFunc2(GetSum, 20, 30)
	fmt.Println("result=", result)

	// 看案例
	a, b := getSumAndSub(1, 2)
	fmt.Println("a=", a)
	fmt.Println("b=", b)
}