package main
import (
	"fmt"
)

/*
	方法实现，键盘接收整数，打印乘法口诀表
*/

type MethodUtils struct {}

func (mu MethodUtils) Print() {
	row := 0
	fmt.Print("请输入一个整数：")
	fmt.Scanln(&row)
	for i := 1; i <= row; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, i * j)
		}
		fmt.Println()
	}
}

func main() {
	
	mu := MethodUtils{}
	mu.Print()
}