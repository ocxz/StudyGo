package main

import "fmt"

// 定义三个全局变量
// var n1, n2, n3 = 100, 200, "jack"

// var n1 = 100
// var n2 = 200
// var n3 = "jack"

// var(
// 	n1 = 100
// 	n2 = 100
// 	n3 = "mary"
// )

// 不支持 n1, n2, n3 := 100, 200, "mary"
func main(){

	// golang 一次性声明多个变量
	// 第一种：var n1, n2, n3 int
	// 第二种：var n1, n2, n3 = 100, "张三", 88
	// 第三种：n1, n2, n3 := 100, "张三丰", 88
	// n1, n2, n3 := 100, "张三丰", 88

	// 第四种
	var(
	n1 = 100
	n2 = 100
	n3 = "mary"
	)
	fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)
}