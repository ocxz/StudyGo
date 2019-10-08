package main
import (
	"fmt"
	"net"
	"time"
	"go_code/14、用户通讯/server/model"
)

/*
	接收流程：
			1、接收客户端发送的信息长度len
			2、再来接受信息本身
			3、判断接受的信息本身长度是否等于len
			4、不等于，就有纠错协议
			5、等于，反序列化成Message结构体变量msg
			6、取出msg.Data，再将其反序列化得到LoginMsg结构体变量loginMsg
			7、取出loginMsg的userName和userPwd
			8、进行数据库匹配，判断是否登录成功
			9、根据比较结果构建信息msg
			10、发送给客户端
*/


// 处理与客户端的通讯
func process(conn net.Conn) {

	// 延时关闭conn
	defer conn.Close()
	// 循环读客户端发送的信息，创建一个处理器，来处理信息
	err := CreateProcessor(conn).processMsg()
	if err != nil {
		fmt.Println("服务器处理消息失败，失败原因：", err)
		return
	}
}

func main() {

	// 当服务器启动时，我们就去初始化redis的连接池
	initPool("localhost:6379", 16, 0, 300 * time.Second)
	model.InitUserDao(pool)
	// 服务器在8889监听
	fmt.Println("服务器在8889监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("服务器监听打开失败，失败原因：", err)
		return
	}

	// 监听成功，等待客户端链接服务器	
	for {
		fmt.Println("等待客户端的连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("客户端连接失败")
		}

		// 拿取到一个与客户端间的conn，开一个线程，为此客户端服务
		go process(conn)

	}
}