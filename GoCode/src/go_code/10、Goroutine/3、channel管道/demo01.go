package main
import (
	"fmt"
	"time"
)

// 计算1-200的各个数的阶乘，结果放入map中，需要使用goroutine完成
// 思路，1、编写一个函数，来计算一个数的阶乘，并放入到map中
// 		2、启动多个协程，将结果放入到map中
//      3、map应该是全局的

var (
	myMap = map[int]int{}
)

// 计算一个数的阶乘n!，并放入到map中
func test(n int) {
	res := 1
	for i :=1; i <= n; i++ {
		res *= i
	}
	myMap[n] = res  // 并发写入问题
}
func main() {

	// 我们这里开启多个线程，完成此任务
	// 这里很丧心病狂的开启了200个协程
	for i := 0; i < 200; i++ {
		go test(i)
	}

	// 主线程等3秒，然后执行下面代码结束程序
	time.Sleep(3 * time.Second)  // 会有等待问题，不知道等多久合适

	// 这里我们输出结果
	for k, v := range myMap {
		fmt.Println("map[%v] = %v\n", k, v)
	}
}