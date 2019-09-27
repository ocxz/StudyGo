package main
import (
	"fmt"
)

func main() {

	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200

	close(intChan)  // 关闭管道

	// 这时不能再向管道中写入数据
	// 会报 panic: send on closed channel 异常
	// intChan <- 300

	fmt.Println("okook~~")

	// 管道关闭后，读取数据是ok的
	n1 := <- intChan
	fmt.Println("intChan关闭后，读取的数据，值为：", n1)

}