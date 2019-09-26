package main
import (
	"fmt"
	"flag"
)

func main() {

	// 定义几个变量，用于接收命令行的参数值
	user := ""
	pwd := ""
	host := ""
	port := 0

	// 接收用户命令行中输入的-u后面的参数值
	// 四个参数：接收、key、默认值、说明
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认3306")

	// 必须调用转换方法
	flag.Parse()

	// 输出结果
	fmt.Printf("用户名=%v 密码=%v 主机号=%v 端口=%v\n", user, pwd, host, port)
}