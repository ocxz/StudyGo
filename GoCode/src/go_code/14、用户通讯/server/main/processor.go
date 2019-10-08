package main
import (
	"fmt"
	"net"
	"go_code/14、用户通讯/common/message"
	"go_code/14、用户通讯/common/utils"
	"go_code/14、用户通讯/server/process"
	"errors"
	"io"
)

type Processor struct {
	Conn net.Conn
}
func CreateProcessor(conn net.Conn) *Processor {
	return &Processor { Conn : conn }
}

// 编写一个处理信息的函数，其功能是根据信息包的种类不同，调用不同的方法解析包，返回结果
func (this *Processor) serverProcessMsg() (err error) {

	// 从conn中，读取一条信息
	msg, err := utils.CreateTransfer(this.Conn).ReadPkg()
	if err != nil {
		if err == io.EOF {
			err = errors.New(fmt.Sprintf("对方关闭了连接：%v", err))
			return
		}
		err = errors.New(fmt.Sprintf("处理信息失败，失败原因：%v", err))
		return
	}

	// 判断信息的类型，进行相应的处理
	switch msg.Type{
		case message.LoginMsgType:     // 登录信息
			// 处理登录
			err = processors.CreateUserProcess(this.Conn).ServerProcessLogin(&msg)
		case message.RegisterMsgType:
			// 处理注册
		default:
			err = errors.New("读取信息失败，失败原因：信息类型不存在，无法处理")
	}
	return
}

// 循环处理与消息
func (this *Processor) processMsg()(err error) {
	for {	
		// 处理信息，接受信息，写回结果
		err = this.serverProcessMsg()
		if err != nil {
			err = errors.New(fmt.Sprintf("服务器处理消息失败，%v", err))
			return
		}
	}
}