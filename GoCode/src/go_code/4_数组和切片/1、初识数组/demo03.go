package main
import (
	"fmt"
)

func main() {

	// 4中定义且初始化数组的方式
	// 第一种：定义数组变量、给数组变量赋值
	var array1 [3]int = [3]int { 1, 2, 3}
	fmt.Println(array1)

	// 第二种：使用类型推导
	var array2 = [3]int {4, 5, 6}
	fmt.Println(array2)

	// 第三种：使用系统推断数组长度
	var array3 = [...]int {1, 2, 3, 4, 5, 6}
	fmt.Println(array3)

	// 第四种，指定下标
	var array4 = [3]string {1:"tom", 2:"marray", 0:"jack" }
	fmt.Println(array4)

	// 第五种，推断长度、指定下标
	var array5 = [...]string {3:"jack", 0:"tom"}
	fmt.Println(array5)
}