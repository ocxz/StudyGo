package main
import (
	"fmt"
	"os"
	"go_code/14、用户通讯/client/process"
)

// 定义用户名和密码
var (
	userName string =  ""
	userPwd string = ""
)

func main() {

	// 接受用户的选择
	var key int
	// 判断是否继续显示菜单
	// loop := true

	for {

		fmt.Println("---------------------欢迎来到多人聊天系统---------------------") 
		fmt.Println("                      1 登录聊天室")
		fmt.Println("                      2 注册用户")
		fmt.Println("                      3 退出系统")
		fmt.Print("  请选择（1-3）：")
		fmt.Scanln(&key)

		switch key {
		case 1 :
			fmt.Println("\n---------------------登录系统---------------------") 
			fmt.Print("请输入用户名：")
			fmt.Scanln(&userName)
			fmt.Print("请输入密码：")
			fmt.Scanln(&userPwd)
			process.CreateUserProcess().Login(userName, userPwd)
		case 2 :
			fmt.Println("\n---------------------注册用户---------------------")
			// loop = false
		case 3 :
			fmt.Println("\n---------------------退出系统---------------------")
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入！")
		}
	}
}