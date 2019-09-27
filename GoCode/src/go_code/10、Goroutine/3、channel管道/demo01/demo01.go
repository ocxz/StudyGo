package main
import (
	"fmt"
	"time"
	"sync"
)

// 计算1-200的各个数的阶乘，结果放入map中，需要使用goroutine完成
// 思路，1、编写一个函数，来计算一个数的阶乘，并放入到map中
// 		2、启动多个协程，将结果放入到map中
//      3、map应该是全局的

var (
	myMap = map[int]int{}
	// 声明一个全局的互斥锁
	// lock 是一个全局的互斥锁，sync是一个包：全称synchornized同步
	lock sync.Mutex
)

// 计算一个数的阶乘n!，并放入到map中
func test(n int) {
	res := 1
	for i :=1; i <= n; i++ {
		res *= i
	}
	// 修改全局变量前，加锁
	lock.Lock()
	myMap[n] = res  // 并发写入问题
	// 修改完成后，解锁
	lock.Unlock()

}
func main() {

	// 我们这里开启多个线程，完成此任务
	// 这里很丧心病狂的开启了200个协程
	for i := 0; i < 20; i++ {
		go test(i)
	}

	// 主线程等3秒，然后执行下面代码结束程序
	time.Sleep(3 * time.Second)  // 会有等待问题，不知道等多久合适

	// 这里我们输出结果
	// 读取全局变量前，加入排他锁
	lock.Lock()
	for k, v := range myMap {
		fmt.Printf("map[%v] = %v\n", k + 1, v)
	}
	// 读取完毕后，解锁
	lock.Unlock()
}