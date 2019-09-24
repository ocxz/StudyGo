package main
import (
	"fmt"
)

/*
	要求：编写方法，提供两个参数，方法中打印m*n的矩形
*/
type MethodUtils struct {}

func (mu MethodUtils) PrintRect(row int, col int) {
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {

	mu := MethodUtils{}
	mu.PrintRect(10, 10)
}