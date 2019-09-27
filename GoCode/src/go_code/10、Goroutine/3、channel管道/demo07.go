package main
import (
	"fmt"
	_ "time"
)

// 完成gorountine和channel协同工作的案例
// 要求：1、开启一个writeData协程，向管道intChan中写入50个整数
//      2、开启一个readData协程，从管道intChan中读取writeData写入的数据
//      3、注意writeData和readData操作的是同一个管道
//      4、主线程需要等待writeData和readData协程都完成工作后才能退出

// 思路：需要借助一个管道，用来阻塞主线程，之前管道中没有内容，但也并未关闭，主线程线程就会一直尝试从该管道读取，产生死锁阻塞
// 当readData协程读50个数据后，即wirteDat和readData完成工作后，就会向第二个管道写入一个标识（可以为true)，并且关闭该管道
// 主线程从管道中不断读取，读到最后，因为该管道已经关闭，所有不再阻塞，读取完成后就程序往下走，乃至退出
func writeData(intChan chan int) {

	for i := 1; i <= 50; i++ {
		// 放入数据
		intChan <- i
		fmt.Println("writeData写入了一个数据，数据 = ", i)
	}
	// 写完后，直接关闭管道，不再往里面写数据了，由写的地方决定管道关闭
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {

	for v := range intChan {
		fmt.Println("readData读到一个数据了，数据 = ", v)
	}
	exitChan <- true
	close(exitChan)
}

func main() {

	// 创建两个管道
	intChan := make(chan int, 5000)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)
	// 用来阻塞主线程
	for v := range exitChan {
		if v {
			fmt.Println("任务完成")
		}
	}
}