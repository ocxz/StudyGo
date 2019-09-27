package main
import (
	"fmt"
)

func main() {

	// 遍历管道
	intChan := make(chan int, 100)

	// 放入100个数据到管道
	for i := 1; i <= 100; i++ {
		intChan <- i * 2
	}

	// 遍历，遍历管道时，不能使用普通的for循环，管道里面没有下标
	// 若没有关闭，就算全部取出，还会继续往下遍历，试图从管道中拿取数据，此时再取就会报错
	// 如果开始遍历时，就将管道关闭，则不会出现死锁异常
	close(intChan) // 关闭管道
	for v := range intChan {
		fmt.Println("v = ", v)
	}
}