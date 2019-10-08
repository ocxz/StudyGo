package process
import (
	"fmt"
	"os"
	"net"
	"go_code/14、用户通讯/common/utils"

)

// 显示登录成功的界面
func ShowMenu() {

	for {
		fmt.Println("----------恭喜XXX登录成功----------")
		fmt.Println("----------1. 显示在线列表----------")
		fmt.Println("----------2. 发送信息----------")
		fmt.Println("----------3. 信息列表----------")
		fmt.Println("----------4. 退出系统----------")
		var key int
		fmt.Print("请选择（1-4）：")
		fmt.Scanln(&key)
		switch key {
			case 1:
				fmt.Println("显示用户在线列表")
			case 2:
				fmt.Println("发送信息")
			case 3:
				fmt.Println("信息列表")
			case 4:
				fmt.Println("退出系统")
				os.Exit(0)
			default:
				fmt.Println("输入不正确，请重新输入！")
		}
	}
}

// 和服务器端保持通讯
func serverProcessMsg(conn net.Conn) {

	// 创建一个transfer实例，不断的读取服务器发送过来的信息
	tf := utils.CreateTransfer(conn)
	for {
		fmt.Println("客户端正在等待读取服务器发送过来的消息")
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("底下读取服务器信息失败，失败原因：", err)
			return
		}
		// 读取后，下一步处理
		fmt.Println("msg = ", msg)
	}
}