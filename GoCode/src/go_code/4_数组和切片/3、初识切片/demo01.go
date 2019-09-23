package main
import (
	"fmt"
)

func main() {

	// 切片的基本使用
	attr := [...]int {1, 2, 3, 4, 5, 6}

	// 切片的获取
	// slice 表示切片名
	// attr[1:3] 表示数组1：3的元素引用，不包括attr[3]
	slice := attr[1:3]
	fmt.Println(attr)
	fmt.Println("数组attr[1:3]的切片内容：", slice)
	fmt.Println("数组attr[1:3]的切片长度：", len(slice))
	// 切片的容量，表示可放置的元素个数，是可以动态变化的
	fmt.Println("数组attr[1:3]的切片容量：", cap(slice))
	

}