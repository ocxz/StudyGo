package main

import (
	"fmt"
)
func PrintRhombus(side int, site int) {
		// 打印空心菱形
	/* 目标：
			*        
		  *   *
		*       *		 
		  *   *
		    *
	*/

	// 第一步：打印边长为3的实心菱形，共3*2 - 1层
	/*
			*           第一层：1个 空格2 
		  * * *         第二层：3个 空格1
		* * * * *	    第三层：5个 空格0	   边长数
		  * * *         第四层：3个 空格1
		    *           第五层：1个 空格2
	*/

	// 第二步：先打印上半层  空格数：边长减去层数  每层*数：2*层数 - 1
	// 第三步：再打印下半层  空格数：层数减去边长数   每层*数：边层所在*数 - 层数边长数之差*2
	// 第四步：镂空菱形，每一层除第一个和最后一个打印*，其余打印空格
	// 第五步：提取变量，控制台输入

	for i := 1; i<= side * 2 - 1 ; i++ {    // 控制层数，打印3*2 - 1 层

		if i <= side {   // 先打上半层
			// 打印空格，每层空格数，等于边长数减去层数
			for j := 1; j <= side - i + site; j++ {   
				fmt.Print("  ")
			}
			// 打印* 每层*个数 等于 层数*2 - 1 
			for j := 1 ; j <= i * 2 - 1; j++ {
				if j == 1 || j == i * 2 - 1 {
					fmt.Print(" * ")
				} else {
					fmt.Print("  ")
				}
				
			}
			fmt.Println()
		}

		if i > side {   // 先打下半层
			// 打印空格，每层空格数，等于层数减去边长数
			for j := 1; j <= i - side + site; j++ {   
				fmt.Print("  ")
			}
			// 打印* 每层*个数 等于 最大长度减去 层数-边长数 * 2
			// 最大数：2*side-1 - (i-side)*2
			for j := 1; j <= 2 * side - 1 - (i-side)*2; j++ {
				if j == 1 || j == 2 * side - 1 - (i-side)*2 {
					fmt.Print(" * ")
				} else {
					fmt.Print("  ")
				}
			}
			fmt.Println()
		}
	}
}
func main() {

	PrintRhombus(10, 0)
	PrintRhombus(10, -5)

}