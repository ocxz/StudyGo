package main
import (
	"fmt"
)

func main() {

	// 数组的内存分析
	// 通过地址的方式，操作数组的各个元素
	var a [5]int
	
	// for i := 0; i <= 5; i++ {
	// 	pert := &a + 8 * i    // 取第i个元素的地址
	// 	*pert = i + 1
	// }

	// p := &a
	fmt.Printf("%v \n", &a[0])
	fmt.Printf("%v \n", &a[1])

}