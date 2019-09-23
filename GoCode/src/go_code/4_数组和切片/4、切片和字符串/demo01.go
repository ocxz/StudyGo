package main
import (
	"fmt"
)

func main() {

	// 利用切片改变字符串
	str := "hello world"
	slice := str[:]
	slice[6] = ""
	fmt.Println(str)
}