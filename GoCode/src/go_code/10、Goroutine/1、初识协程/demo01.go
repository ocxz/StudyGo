package main
import (
	"fmt"
	"time"
)

// 编写一个函数，每隔输出一次，协程说你好
func test(say string, sleepTime time.Duration) {
	for i := 1; i <= 10; i++ {
		fmt.Printf(say + " %v次\n", i)
		// 休息1秒
		time.Sleep(sleepTime)
	}
}

func main() {

	// 主线程中开启一个goroutine协程，每秒输出一次，协程说你好
	// 主线程也每秒输出一次，主线程说你好
	// 主线程输出10次后退出程序
	// 要求主线程和协程同时运行
	go test("协程说你好", time.Second)   // 开启了一个协程
	test("主线程说你好", time.Second)
}