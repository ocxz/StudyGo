package main

import (
	"fmt"
)

// 演示golang指针类型的问题  练习
func main() {

	// 写一个程序，获取一个int变量num的地址，并显示到终端
	// 将num的地址赋给ptr，并通过ptr去修改num的值

	var num int = 10
	var ptr *int = &num
	fmt.Println("ptr指针的值=", ptr)

	// 通过ptr修改num的值
	*ptr = 100
	fmt.Println("修改后，num的值=", num)
}
