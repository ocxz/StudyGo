package main
import (
	"fmt"
)

func main() {

	// 声明一个二维数组
	// 可以看作为 4个[6]int的元素数组
	var arr [4][6]int

	// 初始值为0
	arr[1][2] = 1
	// fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	arr2 := [...][3]int { {1, 2, 3}, {1, 2, 3} }
	fmt.Println(arr2)
}