package main

import (
	"fmt"
)

func main() {

	// 输出
	for i :=1; i <= 10; i++ {
		fmt.Println("你好，世界")
	}

	// for 循环的第二种写法，将循环变量初始化和迭代写到其他位置
	j := 1  // 循环变量初始化
	for j <= 10 {
		fmt.Println("hello world", j)
		j++   // 循环变量迭代
	}

	// for 循环第三种用法 通常需要配合break语句来使用
	k := 1
	for {
		if k <= 10 {
			fmt.Println("ok~", k)
		} else {
			break;   // 跳出这一层的for循环
		}
		k++
	}

	// 字符串遍历方式一  传统方式
	var str string = "hello world!"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])   // 使用到了字符串下标
	}
	fmt.Println()

	// 使用for-range方式 遍历字符串
	str = "abc-ok"
	for index, value := range str {   // for range 取出str中的索引以及索引值
		fmt.Printf("index=%d, val=%c \n", index, value)
	}

	// 切片，解决传统方式遍历中文字符串出现乱码问题
	var str2 string = "hello world! 北京"
	str2_rune := []rune(str2)  // 将str2转为[]rune 切片类型
	for i := 0; i < len(str2_rune); i++ {
		fmt.Printf("%c ", str2_rune[i])   // 使用到了字符串下标
	}
	fmt.Println()
}