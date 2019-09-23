package main
import (
	"fmt"
)

func BinaryFind(arr []int, findValue int) int {

	leftIndex := 0
	rightIndex := len(arr) - 1
	come:
	for {

		midIndex := (rightIndex + leftIndex) / 2
		switch {
			case rightIndex < leftIndex:
				return -1
			case findValue < arr[midIndex]:
				rightIndex = midIndex - 1
				continue come
			case findValue > arr[midIndex]:
				leftIndex = midIndex + 1
				continue come
			default:
				return midIndex
		}
	}
}
func main() {

	arr := []int {1, 2, 3, 4, 5, 6}
	for {
		findValue := 0
		fmt.Scanln(&findValue)
		findIndex := BinaryFind(arr, findValue)
		if findIndex < 0 {
			fmt.Printf("没找到%v \n",findIndex)
		} else {
			fmt.Printf("找到了，下标为%v，值为%v \n", findIndex, findValue)
		}
	}
}