package main

import (
	"fmt"
)

func main() {

	i := 50
	// 二进制输出
	fmt.Printf("%v的二进制输出是：%b \n", i, i)

	// 八进制：0-7 以字母0开头表示
	var j int = 011   // 对应十进制中的9
	fmt.Println("j=", j)

	// 十六进制：0-9A-F 以0x或0X开头
	var k int = 0x11  // 对应十进制中的17
	fmt.Println("k=", k)
}