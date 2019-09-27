package main
import (
	"fmt"
)

func main() {

	// 演示一下管道的使用
	// 1、创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	// 2、看看intChan到底是个啥
	fmt.Printf("intChan = %p ，intChan本身的地址%p\n", intChan, &intChan)

	// 3、向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num

	// 4、看看管道的长度和cap(容量)，管道不会自动增长，超出容量就报错
	fmt.Printf("intChan的长度是：%v，容量是%v\n", len(intChan), cap(intChan))

	// 5、从管道中取出数据
	num2 := <- intChan
	fmt.Println("num2 = ", num2)
	fmt.Printf("取出num2后，intChan的长度是：%v，容量是%v\n", len(intChan), cap(intChan))

}