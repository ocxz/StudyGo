package main
import (
	"fmt"
)

func main() {

	// 使用select，可以解决管道取数据的阻塞问题
	// 1、定义一个管道 10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	// 2、定义一个管道 5个数据string
	strChan := make(chan string, 5)
	for i :=0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%v", i)
	}

	// 传统方法在，遍历管道时，如果不关闭，会阻塞而导致死锁
	// 有时候确实很难确定什么时候关闭管道，这时候就可以使用
	// select 方式来解决
	// 每次都会尝试从执行case，取数据，没有case匹配则说明管道中都没有数据了
	for {
		select {
			// 注意：这里即使管道一直没有关闭，也不会一直阻塞而死锁，会自动到下一个case匹配
			case v := <-intChan:  
				fmt.Println("从intChan管道中，读取到了数据", v)
			case v := <- strChan:
				fmt.Println("从strChan管道中，取到了数据", v)
			default:
				// 程序员可以加入逻辑
				fmt.Println("都取不到了，不玩了")
				return
		}
	}
}