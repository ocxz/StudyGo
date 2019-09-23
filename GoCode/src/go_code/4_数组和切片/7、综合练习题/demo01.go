package main
import (
	"fmt"
	"math/rand"
	"time"
)

/*
	随机生成10个整数
	倒序打印
	求平均值、最大值、最大值下标、查找里面是否有55
*/

func main() {

	// 生成10个随机整数
	var arr [10]int
	for i, _ := range arr {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(10 * time.Microsecond)
		arr[i] = rand.Intn(100) + 1
	}
	fmt.Println(arr)

	// 倒序打印
	for i := len(arr) -1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// 求最大值，以及最大值下标
	maxIndex := 0
	maxValue := 0
	for i := 0; i < len(arr); i++ {
		if maxValue < arr[i]{
			maxValue = arr[i]
			maxIndex = i
		}
	}

	// 查找是否有55，使用二分法
	// 第一步：排序
	fmt.Printf("最大值为%v，最大值的下标为%v \n", maxValue, maxIndex)

	for i := 0; i < len(arr) - 1; i++ {
		for j :=0; j < len(arr) - 1 - i; j++ {
			if arr[j] < arr [j+1]{
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)
	// 第二步查找
	leftInd
}