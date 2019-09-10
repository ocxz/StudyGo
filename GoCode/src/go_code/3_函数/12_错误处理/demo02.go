package main
import (
	"fmt"
	"errors"
)

// 模拟读取配置文件的信息，文件名：config.ini
// 如果文件名不准确，则抛出自定义错误，终止执行
func readConf(name string) (err error) {
	if name == "config.init" {
		// 读取
		return nil   // 没有错误
	}else {
		// 否则返回自定义错误
		return errors.New("读取配置文件错误，配置文件名不正确")
	}
}

// 调用readConf，进行配置文件的读取，处理异常
func test02() {

	err := readConf("config.ini")
	if err == nil {
		// 没有错误，执行代码
	} else {
		// 有错误，捕获异常，终止执行
		panic(err)   // 打印错误信息，且终止程序
	}
	fmt.Println("test02继续执行")
}

func main() {

	// 自定义错误处理
	test02()
	fmt.Println("main下面的代码")
}