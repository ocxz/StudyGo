package main
import (
	"net"
	""
)

/*
	用来定义服务器，结构体
	字段包括：协议、ip、端口号、在线客户
	方法：接收信息、发送信息、添加客户在线、删除退休客户
	函数：创建服务器，并开启服务
*/

type Server struct {
	Pcl string
	IP string
	Port int
	Users []*User
}

// 创建一个服务器
func NewServer(pcl string, Ip string, port int) *Server {

	return &Server {
		Pcl : pcl,
		IP : Ip,
		Port : port,
	} 
}

// 判断一个用户是否在线，负一就是不在线
func (server *Server)Online(u *User) int {

	for k, v := range server.Users {
		if u == v {
			return k
		}
	}
	return -1
}

// 添加一个上线用户，如果该用户已上线，则添加失败
func (server *Server)AddUser(u *User) bool {

	if server.Online(u) >= 0 {   // 用户已上线，添加失败
		return false
	}
	server.Users = append(server.Users, u)
	return true
}

// 删除一个下线用户
func (server *Server)DelUser(u *User) bool {
	
	index := server.Online(u)
	if index < 0 {   // 如果用户不在线，则删除失败
		return false
	}

	// 删除下标为index的用户
	server.Users = append(server.Users[:index], server.Users[index+1:]...)
}

// 服务器向客户端发送信息，仅仅发送一次
func (server *Server)SendMsg(msg string, users...*Users) {

	usersNum := len(users)
	if
}
