package main
import (
	"fmt"
	"time"
	"strconv"
)

var errTime int;
// 测试goLang异常捕获和处理
func test() {
	// 使用defer + recover 来捕获处理异常
	defer func() {
		err := recover()   // recover() 内置函数，可捕获并返回异常
		if err != nil {    // 说明捕获到错误
			fmt.Println("err=", err)
			errTime++
			// 这里可以处理错误
		}
	}()
	num1,num2 := 1, 0
	result := num1 / num2
	fmt.Println("result=", result)
}

func main() {

	// 测试golang遇到异常代码
	for ;errTime <10; {
		test()
		fmt.Println("main函数下面的代码" + strconv.Itoa(errTime ))
		time.Sleep(time.Second)
	}
}