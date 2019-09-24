package main
import (
	"fmt"
)

/*
	要求：根据行、列、字符 打印字符组成的矩形
*/

type MethodUtils struct{}

func (mu *MethodUtils) PrintRect(row int, col int, char string) {
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			fmt.Print(char+" ")
		}
		fmt.Println()
	}
}

func main() {

	mu := MethodUtils{}
	mu.PrintRect(7, 20, "@")
}
