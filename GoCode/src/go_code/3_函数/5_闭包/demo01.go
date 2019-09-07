package main

import (
	"fmt"
	"strings"
)

// 闭包
// AddUpper是一个函数，返回的类型是func(int) int
// 闭包的说明：返回的是一个匿名函数，但匿名函数引用到了函数外的变量n，因此这个匿名函数与n变量形成一个整体，构成闭包
// 可以理解为：闭包是一个类，n是字段，函数于n构成一个整体，也就是闭包
// 当我们反复调用闭包中的函数时，因为n只初始化一次，所以n的值会得以保存
// 闭包的关键：分析出返回函数所引用到的变量，因为函数和他引用的变量构成闭包
func AddUpper() func (int) int {

	// 下面的代码构成一个闭包
	var n int = 10
	var str string = "hello"
	return func (x int) int {
		n += x
		str += string(int('0')+x )
		fmt.Println("str=", str)
		return n
	}
}

// 编写一个makeSuffix(suffix, string)函数，可接受一个文件后缀名（如.jpg）返回一个相应的闭包
// 调用该闭包，传入文件名，若没有指定后面名则添加，否则返回即可
// strings.HasSuffix(name, suffix)  可用判断某个字符串是否具有指定的后缀
// 返回匿名函数与引用的suffix变量构成闭包
func makeSuffix(suffix string) func(string) string {

	// 如果没有指定后缀，则添加 否则返回
	return func(name string) string {
		if !strings.HasSuffix(name, suffix){
			return name + suffix
		} else {
			return name
		}
	}
}

func main() {

	// 累加器 使用闭包
	f := AddUpper()
	fmt.Println(f(1))   // 调用一次 返回匿名函数， n：10+1=11 保存了起来
	fmt.Println(f(2))   // 第二次调用返回匿名函数  n：11+2=13 保存了起来
	fmt.Println(f(3))   // 第三次调用返回匿名函数  n：13+3=16 保存了起来

	f2 := makeSuffix(".jpg")
	fmt.Println(f2("臭小子"))
	fmt.Println(f2("臭小子.jpg"))
	fmt.Println(f2("臭小子.jsp"))
}