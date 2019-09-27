package main
import (
	"fmt"
)

/*
	要求：
		1、启动一个协程，将1-2000的数放入到numChan中
		2、启动8个协程，从numChan取出数据，并计算n!，将结果存放到resChan中
		3、最后8个协程协同完成工作后，遍历resChan，显示结果res[1] = 1
		4、注意考虑resChan chan int 是否合适
		5、答案是否定的，因为8个协程拿数据是排队的，但是计算时间不一定一样，因此写resChan时可能顺序会有问题
	分析：
		1、需要3个管道，numChan、reChan、exitChan
		2、几个协程没有关系，因为numChan由写协程管控，其它八个协程共同线程安全访问numChan，当numChan
		   关闭，且某个协程拿取完最后一个数据，就会通过exitChan通知主线程，读取数据完毕，主线程解开阻塞
		   继续执行，直到程序结束
*/

// 向numChan写数据
func writeData(numChan chan int, n int) {

	for i := 1; i <= n; i++ {
		numChan <- i
	}
	// 写完后，关闭管道
	close(numChan)
}

// 计算阶乘
func factorial(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

// 读取numChan中的数据，计算阶乘，放入resChan中，通知主线程执行完毕
func readData(numChan chan int, resChan chan map[int]uint64, finishChan chan bool) {

	for v := range numChan {
		resMap := <- resChan   // 将map[int]uint64取出
		resMap[v] = uint64(factorial(v))  // 给map添加一个值
		resChan <- resMap  // 将map重新放到管道中
	}

	// 读取完成，关闭resChan，通知主线程执行完毕
	fmt.Println("有一个readData，取不到数据，退出了")
	finishChan <- true
}

func main() {

	n := 100  // 定义需要多少个
	numChan := make(chan int, 20)
	resChan := make(chan map[int]uint64, 1)
	finishChan := make(chan bool, 8)
	// exitChan := make(chan bool, 1)
	resChan <- map[int]uint64 { 1:1}   // 需要放入一个初始值

	go writeData(numChan, n)  // 开启一个写协程

	// 开启八个读协程
	for i := 0; i < 8; i++ {
		go readData(numChan, resChan, finishChan)
	}

	// 这里主线程做出处理，关键时从exitChan取出8个true，就证明readData八个协程工作全部完成

	// 当从exitChan取出来8个结果，就可以关闭resChan
	// 开启一个协程，用来关闭resChan
	// go exit(finishChan, resChan, exitChan)
	var resMap map[int]uint64

	// 使用匿名协程，用来专门管理什么时候结束
	// 思想，就是从finishChan丢八次数据，否则就一直阻塞，当丢够了八次数据后，关闭resChan、finishChan，将结果取出
	go func() {
		for i := 0; i < 8; i++ {
			<- finishChan
		}
		close(resChan)
		// close(finishChan)
		resMap = <- resChan
	}()

	// 阻塞判断resMap是否由结果，有则说明读写操作完成，把结果返回了
	for {
		if len(resMap) !=0 {
			break
		}
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("res[%v] = %v \n", i, resMap[i])
	}

}
