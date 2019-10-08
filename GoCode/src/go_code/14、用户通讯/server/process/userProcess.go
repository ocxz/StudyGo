package processors
import (
	"fmt"
	"net"
	"go_code/14、用户通讯/common/message"
	"go_code/14、用户通讯/common/utils"
	"go_code/14、用户通讯/server/model"
	"encoding/json"
	"errors"
)
type UserProcess struct {
	Conn net.Conn
}
func CreateUserProcess(conn net.Conn) *UserProcess {
	return &UserProcess{ Conn : conn }
}

// 编写一个函数，专门处理登录信息
func (this *UserProcess) ServerProcessLogin(msg *message.Message) (err error) {

	// 从msg中取出data，并反序列化成LoginMsg类型
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		err = errors.New(fmt.Sprintf("消息体反序列化失败，失败原因：%v", err))
		return
	}
	// 如果用户userName = 123456，userPwd = 123456 则表示用户合法
	// 声明一个响应信息体
	var resMsg message.Message
	resMsg.Type = message.ResultMsgType
	// 声明一个登录结果信息体
	var result message.ResultMsg
	
	// 去redis数据库，完成数据校验
	user, err := model.UserDao.Login(loginMsg.UserName, loginMsg.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			result.Code = 500
			result.Error = "用户不存在"
		}
		if err == model.ERROR_USER_PWD {
			result = 400
			result.Error = "用户密码不正确"
		}
	} else {
		result.Code = 200
		fmt.Println(user)
	}

	// 将结果序列化
	rm, err := json.Marshal(result)
	if err != nil {
		err = errors.New(fmt.Sprintf("登录返回结果序列化失败，失败原因：%v", err))
		return
	}
	// 将结果序列化串赋值给resMsg响应信息结构体
	resMsg.Data = string(rm)

	// 将响应信息结构体序列化，准备发送
	data, err := json.Marshal(resMsg)
	if err != nil {
		err = errors.New(fmt.Sprintf("响应结构体序列化失败，失败原因：%v", err))
		return
	}
	// 发送信息给客户端
	err = utils.CreateTransfer(this.Conn).WritePkg(data)
	return
}
