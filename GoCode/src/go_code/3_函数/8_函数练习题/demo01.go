package main

import (
	"fmt"
)

// 输入层数，打印金字塔
func PrintJiziTa(level int) {
 
	// 打印金字塔
	// 打印空心金字塔 除最后一行，只打印第一个最后一个
	// 最后一行空格打印
	for i := 1; i <= level; i++ {   // 层数

		// 打印空格 如3层 依次是：2、1、0 所以空格数为level-i
		for k := 0; k < level - i; k++{
			fmt.Print(" ")
		}

		// 打印*每层打印 1 3 5 7 也就是2i- 1 个*
		for j := 1 ; j <= 2 * i - 1; j++{
			if i == level {
				if j % 2 == 0 {
					fmt.Print(" ")
				}else{
					fmt.Print("*")
				}
				continue
			}
			if j == 1 || j == 2 * i - 1{
				fmt.Print("*")
			}else {
				fmt.Print(" ")
			}	
		}
		// 换行打印
		fmt.Println()
	}
}

// 输入层数，打印口诀表
func PrintKouJue(level int) {

	// 第一层1个 1*1  第二层2个 1*1  1*2  第一个数是列数 第二个数是层数
	for i := 1; i <= level; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, j * i)
		}
		fmt.Println()
	}
}

func main() {
	PrintJiziTa(4)
	PrintKouJue(4)
}