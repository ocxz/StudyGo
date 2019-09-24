package main
import (
	"fmt"
)

/*
	要求：编写一个方法，给定一个二维数组，将其转置
*/

type MethodUtils struct{}

func (mu *MethodUtils) tdArrayConv(arr [3][3]int) [3][3]int {
	var tempArray [3][3]int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			tempArray[j][i] = arr[i][j]
		}
	}
	return tempArray
}

func main() {

	mu := MethodUtils{}
	arr := [3][3]int {{1, 2, 3}, {4, 5, 6}, {7, 8, 8}}
	fmt.Println(arr)
	arr = mu.tdArrayConv(arr)
	fmt.Println(arr)
}