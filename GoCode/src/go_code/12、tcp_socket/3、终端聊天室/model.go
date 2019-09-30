package main
import (
	"net"
	"encoding/json"
)
const (
	SPcl = "tcp",
	SIp = "192.168.0.195",   // 服务器的Ip
	SPort = "8888"
)
// 定义一个通讯用户
type User struct {
	Name string `json:"name"`    // 用户名，唯一的
	Pwd string	`json:"pwd"`      // 密码
	Status bool `json:"status"`    // 在线状态
	Friends map[string]string `json:"friends"`   // 好友，唯一用户、第二个是昵称
	conn net.Conn,   // 用来存储，本次连接的一个状态
}

// 创建一个通讯用户
func NewUser(name string, pwd string) *User {

	// 给服务端发送一次请求，表明我要注册
	conn, err := net.Dial(SPcl, SIp + ":" + SPort)
	if err != nil {
		fmt.Println("连接服务器失败，失败原因：", err)
		return nil
	}
	// 构建一个用户json格式数据
	user := &User {Name : name, Pwd : pwd}
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("转换为json格式数据错误，错误原因：", err)
		return nil
	}

	// 将此json串发送给服务器
	if !user.SendMsgBytes(json) {
		return nil
	}

	// 接收服务器传过来的json格式数据
	rdata := mu.ReceiveMsgBytes()
	if rdata == nil {
		return nil
	}
	err = json.Unmarshal(rdata, &user)
	if err != nil {
		fmt.Println("接收数据转化失败，失败原因：", err)
		return nil
	}
	user.conn = conn
	return user
}

// 添加朋友，可指定昵称
func (mu *User) AddFriend(nickName string, user *User) bool {

	name := nickName
	if nickName == "" {
		name = user.Name
	}
	name += "_" +user.IP

	// 判断朋友是否已存在
	_, isExist := mu.Friends[name]

	if isExist {
		return false
	}

	// 添加朋友
	mu.Friends[name] = user
}

// 向服务器发送一条信息
func (mu *User) SendMsg(msg string) bool {

	// 向服务器，发送一条数据
	return mu.SendMsgBytes([]byte(msg))
}

// 向服务器发送一个二进制切片
func (mu *User) SendMsgBytes([]byte) bool {

	_, err := mu.Conn.Write([]byte)
	if err != nil {
		fmt.Println("数据发送给服务器时失败，失败原因：\n", err)
		return false
	}
	return true
}

// 从服务器接收一条信息，统一换成字符串
func (mu *User) ReceiveMsg() string {

	bytes := mu.ReceiveMsgBytes()
	if bytes == nil {
		return ""
	}
	return string(mu.ReceiveMsgBytes())
}

// 接收服务器发送过来的二进制切片
func (mu *User) ReceiveMsgBytes() []byte {

	// 从服务器接收数据
	buf := make([]byte, 1024)
	n, err := mu.Conn.Read(buf)
	if err != nil {
		fmt.Println("从服务器接收数据失败，失败原因：", err)
		return nil
	}
	return buf[0:n]
}