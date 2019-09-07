package utils
import (
	"fmt"
)

var A int = testA()
var B int = testB()
var C int = testC()

func testA() int {
	fmt.Println("初始化utils中的变量A")
	return 100
}

func testB() int {
	fmt.Println("初始化utils中的变量B")
	return 200
}

func testC() int {
	fmt.Println("初始化utils中的变量C")
	return 300
}

func init() {
	fmt.Println("执行了utils中的init函数")
}

func Test() {
	fmt.Println("调用了utils中的普通函数")
}