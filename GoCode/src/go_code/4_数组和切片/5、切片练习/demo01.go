package main
import (
	"fmt"
)
	// 要求：
	// 编写函数，可接收n int
	// 将斐波那契数列放入切片中
	// 示例：attr[0] = 1  attr[1] = 1  attr[2] = 2
func fibonacci(n int) []uint64 {

	fbnSlice := make([]uint64, n)
	switch {
		case n <= 1:
			fbnSlice[0] = 1
			return fbnSlice
		case n <= 2:
			fbnSlice[0] = 1
			fbnSlice[1] = 1
			return fbnSlice
		default:
			fbnSlice[0] = 1
			fbnSlice[1] = 1
			for i :=2; i < n; i++ {
				fbnSlice[i] = fbnSlice[i - 1] + fbnSlice[i - 2]
			}
			return fbnSlice
	}
}
func main() {
	
	slice := fibonacci(5)
	fmt.Println(slice)
}