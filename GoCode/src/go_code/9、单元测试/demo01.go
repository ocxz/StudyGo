package main
import (
	"fmt"
)

func AddUpper(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func main() {

	// 传统方式测试AddUpper()函数是否正确
	res := AddUpper(10)
	fmt.Println(res)
}