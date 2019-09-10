package main
import (
	"fmt"
	"time"
	"strconv"
)

// 拼接字符串字符串十万次，测试方法执行的时间
func test() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i) + "\n"
	}
}

func main() {

	// 拼接字符串字符串十万次，测试方法执行的时间
	// 记录执行开始的时间戳
	fmt.Println("开始执行")
	start := time.Now().Unix()	
	test()

	// 记录执行完毕的时间戳
	end := time.Now().Unix()
	fmt.Println("执行完毕")
	fmt.Printf("拼接10w次字符串所用的时间是：%v秒", end - start)
}