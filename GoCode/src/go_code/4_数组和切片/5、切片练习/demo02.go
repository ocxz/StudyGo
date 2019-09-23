package main
import (
	"fmt"
)

// 冒泡排序
func SortArray(array []int) {

	for i := 0; i < len(array) -1; i++ {   // 比较轮次，一共需要比较数组长度减一轮
		for j :=0; j < len(array) - (i + 1); j++ {   // 每次需要比较数组长度减去轮次
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
}

func main() {
	array := []int {1, 30, 58, 63, 21, 23, 78, 64}
	SortArray(array)
	fmt.Println(array)
}