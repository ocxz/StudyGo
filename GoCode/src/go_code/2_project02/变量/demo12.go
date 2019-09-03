package main

import (
	"fmt"
	"strconv"
)

// 演示golang的基本数据类型与string转换问题
func main() {

	var num1 int = 99
	var num2 float64 = 23.145
	var b bool = true
	var myChar byte = 'h'
	var str string

	// 方式一：使用fmt.Sprintf()方法，进行转换
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type：%T, str=%q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type：%T, str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type：%T, str=%q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type：%T, str=%q\n", str, str)

	// 方式二：使用strconv包中函数

	fmt.Println("\n")
	str = strconv.FormatInt(int64(num1), 10)   // 10表示以十进制方式进行转换
	fmt.Printf("str type：%T, str=%q\n", str, str)

	// Int转为string类型的简写形式
	str = strconv.Itoa(num1)
	fmt.Printf("str type：%T, str=%q\n", str, str)

	// 'f' 表示格式  5 表示小数保留位数  64表示要转换浮点数的位数
	str = strconv.FormatFloat(num2, 'f', 5, 64)
	fmt.Printf("str type：%T, str=%q\n", str, str)

	str = strconv.FormatBool(b)
	fmt.Printf("str type：%T, str=%q\n", str, str)


}
