package main
import (
	"fmt"
	"math"
	"time"
)

/*
	要求：打印1-8000之间的所有素数、1协程写、4协程读、1协程做结束决策
*/

// 判断该素数的方法
// 处理1和本身可整除，也就是说只有2-自己的开根取余没有等于0，则说明是素数
func PrimeNum(n int) bool {
	for i :=2; i <= int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

// 向intChan管道中，写入数据
func writeData(intChan chan int, n int) {

	for i := 1; i <= n; i++ {
		intChan <- i
	}
	close(intChan)
	fmt.Println("writeData写的协程工作完成")
}



// 从intChan读数据，判断是否为质数，加入到resChan中，协程工作完成后，将状态加入到finishChan中
func readData(intChan chan int, resChan chan int, finishChan chan bool) {

	for data := range intChan {
		if PrimeNum(data) {
			resChan <- data
		}
	}

	// 一个readData协程完成工作，将完成工作状态加入到finishChan管道中
	finishChan <- true
	fmt.Println("writeData读的一条协程工作完成")
}

// 同一管控readData协程，当所有协程完成工作后，通知主线程结束操作
func finish(resChan chan int, finishChan chan bool, exitChan chan bool, chanNum int) {

	for i := 0; i < chanNum; i++ {
		<- finishChan
	}

	close(finishChan)
	close(resChan)
	exitChan <- true
	close(exitChan)
}

func main() {

	n := 4  // 要开4条读协程
	num := 8000000  // 求1-8000之间的素数
	intChan := make(chan int, 1000)
	resChan := make(chan int, num / 2)
	finishChan := make(chan bool, n)  // 管控是否读完了
	exitChan := make(chan bool, 1)

	// 记录开始数据 
	start := time.Now().Unix()
	go writeData(intChan, num)  // 开一条写协程
	for i := 0; i < n; i++ {   // 开n条读协程
		go readData(intChan, resChan, finishChan)
	}

	go finish(resChan, finishChan, exitChan, n)  // 开一个控制结束的协程

	// 阻塞
	for v := range exitChan {
		if v {
			break
		}
	}
	// 记录结束时间
	end := time.Now().Unix()
	fmt.Printf("%v条协程，完成工作，耗费的时间是：%v秒", n, end - start)

	// for w := range resChan {
	// 	fmt.Println(w, "是素数")	
	// }


	// for v := range exitChan {
	// 	if v {
	// 		for w := range resChan {
	// 			fmt.Println(w, "是素数")
	// 		}
	// 	}
	// }
}