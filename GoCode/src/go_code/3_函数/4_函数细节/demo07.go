package main

import (
	"fmt"
)

var age = test()
// 为了看到全局变量是先被执行的，我们可以写一个函数作为全局变量返回值
func test() int {
	fmt.Println("全局变量初始化")
	return 90
}

// init 函数
func init(){
	fmt.Println("执行了init函数")
}
func main() {

	fmt.Println("执行了main函数")
}