package main

import (
	"fmt"
)

// 演示golang的字符串类型
func main() {

	// string的基本使用
	var address string = "北京长城 110 Hello world"
	fmt.Println(address)

	// // 字符串不可变，一旦赋值不可变
	// var str = "Hello"

	// 反引号  原生字符
	str2 := "abc\n\"abc\""
	fmt.Println(str2)

	str3 := `
package main

import(
"fmt"
)

// 演示golang的字符串类型
func main(){

// string的基本使用
var address string = "北京长城 110 Hello world"
fmt.Println(address)

// // 字符串不可变，一旦赋值不可变
// var str = "Hello"

// 反引号  原生字符
str2 := "abc\n\"abc\""
fmt.Println(str2)
}
	`

	fmt.Println(str3)

}
