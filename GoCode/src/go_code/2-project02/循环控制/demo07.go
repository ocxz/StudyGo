package main

import (
	"fmt"
)

func main() {

	// 打印乘法口诀表
	for i := 1; i <= 9; i++ {    // 控制九层
		// 控制每层个数，第一层一个，第二层2个
		for j := 1; j <= i; j++ {
			// 打印：列数 * 层数 = 结果
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}