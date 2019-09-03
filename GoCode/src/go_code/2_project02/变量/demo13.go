package main

import (
	"fmt"
	"strconv"
)

// 演示golang的string转换基本数据类型的问题
func main() {
	var str string = "true"
	var str2 string = "123456789"
	var str3 string = "123456.789"

	// var b bool = strconv.ParseBool(str)
	b, _ := strconv.ParseBool(str)   // ParseBool会返回两个值，转换和erro  erro我们可以用_来忽略
	fmt.Printf("b type：%T，b=%v \n", b, b)

	num, _ := strconv.ParseInt(str2, 10, 64)
	fmt.Printf("num type：%T，num=%v \n", num, num)
	

	f, _ := strconv.ParseFloat(str3, 64)
	fmt.Printf("f type：%T，f=%v \n", f, f)

	// 注意：
	var str4 = "hello world"
	n3, _ := strconv.ParseInt(str4, 10, 64)
	fmt.Printf("n3 type：%T，n3=%v \n", n3, n3)
}
