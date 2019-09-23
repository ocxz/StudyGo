package main
import (
	"fmt"
)

func main() {

	// 创建一个byte类型的26个元素的数组，并打印出来
	var arr [26]byte
	for i, _ := range(arr) {
		arr[i] = 'A' + byte(i)
	}

	for _, v := range(arr) {
		fmt.Printf("%c ", v)
	}
}