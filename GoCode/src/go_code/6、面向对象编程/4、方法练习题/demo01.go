package main
import (
	"fmt"
)

// 要求：编写结构体(MethodUtils)，编写方法，不需要参数，方法中打印10 * 8 的矩形，main方法中调此方法

// 定义结构体
type MethodUtils struct {

}

// 定义方法，绑定到结构体中
func (methodUtils MethodUtils) PrintRect() {
	row := 10
	col := 8
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {

	mu := MethodUtils{}
	mu.PrintRect()
}