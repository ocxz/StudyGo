package main
import (
	"fmt"
)

func main() {

	// 演示二维数组的遍历
	arr := [...][3]int { {1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	// 双层for循环遍历
	for i := 0; i < len(arr); i++ {
		for j :=0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j] , " ")
		}
		fmt.Println()
	}
	fmt.Println()

	// for-range遍历
	for _, v := range arr {
		for _, w := range v {
			fmt.Print(w, " ")
		}
		fmt.Println()
	}
}