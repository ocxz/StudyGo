package process
import (
	"fmt"
	"net"
	"encoding/json"
	"go_code/14、用户通讯/common/message"
	"go_code/14、用户通讯/common/utils"
)

type UserProcess struct{

}
func CreateUserProcess() *UserProcess {
	return &UserProcess{}
}

// 写一个函数，完成登录校验
func (this *UserProcess) Login(userName string, userPwd string) (err error) {

	// 下一个就要开始定协议
	// fmt.Printf("userName = %v，userPwd = %v\n", userName, userPwd)
	// return nil
	// 连接到服务器
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("连接服务器失败，失败原因：", err)
		return
	}
	defer conn.Close() // 延时关闭

	// 准备通过conn，发送消息给服务器
	var msg message.Message
	msg.Type = message.LoginMsgType
	// 创建一个LoginMsg 结构体
	var loginMsg message.LoginMsg
	loginMsg.UserName = userName
	loginMsg.UserPwd = userPwd

	// 将loginMsg序列化，赋值给data
	data, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("结构体数据序列化失败，失败原因：", err)
		return
	}
	msg.Data = string(data)
	// 将msg序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("结构体数据序列化失败，失败原因：", err)
		return
	}

	// 将登录信息的序列化数据写给服务器
	err = utils.CreateTransfer(conn).WritePkg(data)
	if err != nil {
		fmt.Println("写包失败，失败原因：", err)
		return
	}

	// 这里需要处理服务的返回的消息
	msg, err = utils.CreateTransfer(conn).ReadPkg()
	if err != nil {
		fmt.Println("读包失败，失败原因：", err)
		return
	}
	// 将msg中的Data反序列化
	var result message.ResultMsg
	err = json.Unmarshal([]byte(msg.Data), &result)
	if err != nil {
		fmt.Println("结果体反序列化失败，失败原因：", err)
		return
	}

	switch result.Code {
	case 200:

		// 开启一个协程用于和服务器的通讯
		// 如果服务器有数据推送给客户端，则接受并显示在客户端
		go serverProcessMsg(conn)

		// 登录成功，显示菜单
		ShowMenu()
	case 500, 400:
		fmt.Println("登录失败，失败原因：", result.Error)
	default:
		fmt.Println("响应码不存在，无法处理")
	}
	return
	// msg.Data = string(data)
}