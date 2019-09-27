package main
import (
	"fmt"
	"time"
)

func sayHello() {

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

func test() {

	// 这里使用defe + recover
	defer func() {

		if err := recover(); err != nil {
			fmt.Println("test()发送错误，错误原因：", err)
		}
	}()
	// 定义了一个map
	var myMap map[int]string
	myMap[0] = "golang"
}

func main() {

	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println("main()说：hello")
	}
}