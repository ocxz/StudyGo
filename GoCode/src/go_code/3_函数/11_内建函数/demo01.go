package main
import (
	"fmt"
)

func main() {

	// new的使用
	num1 := 100
	num3 := 100
	fmt.Printf("num1的类型%T，num1的值=%v，num1的地址=%v \n", num1, num1, &num1)
	fmt.Printf("num3的类型%T，num3的值=%v，num3的地址=%v \n", num3, num3, &num3)

	num2 := new(int)   // *int 即int指针类型
	*num2 = 1000
	fmt.Printf("num2的类型%T，num2的值=%v，num2的地址=%v，num2指向的值%v \n", num2, num2, &num2, *num2)
}