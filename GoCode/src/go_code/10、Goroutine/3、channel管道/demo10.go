package main
import (
	"fmt"
)

func main() {

	// 管道可以声明为只读或者只写
	// 默认情况下，管道是双向的

	// 只写管道
	var chan2 chan<- int
	chan2 = make(chan int, 3)

	chan2 <- 20
	fmt.Println("chan2 = ", chan2)

	// 会报错 invalid operation: <-chan2 (receive from send-only type chan<- int)
	// num := <- chan2
	// fmt.Println("num = ", num)

	// 只读管道
	var chan3 <-chan int
	num2 := <- chan3
	fmt.Println(num2)
}