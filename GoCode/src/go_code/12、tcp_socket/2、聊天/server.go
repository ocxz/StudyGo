package main
import (
	"fmt"
	"net"   // 做网络socket开发时，net包含我们所需的所有方法和函数
	_ "io"
)

// 处理客户端的请求
func process(conn net.Conn, conns []net.Conn) {

	// 这里我们循环接收客户端发送的数据
	defer conn.Close()   // 关闭conn连接
	client := conn.RemoteAddr().String()   // 网络名
	fmt.Printf("服务器在等待客户端%v的输入\n", client)
	for {
		// 创建一个新的切片
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)    // 客户端没有发送数据，就会一直阻塞，如果客户端没有write，那么协程就会阻塞在这里
		if err != nil {
			fmt.Println("服务端读取客户端发送的数据失败，失败原因：", err)
			return
		}
		msg := string(buf[:n])   // 客户端发过来的信息
		if msg == "exit" {
			fmt.Printf("客户端%v结束了会话\n", client)
			return
		}

		msg = client + "说：" + msg
		fmt.Println(msg)
		for _, v := range conns {
			_, err := v.Write([]byte(msg))
			if err != nil {
				fmt.Println("信息转发给%v失败，失败原因：", err)
			}
		}
	}
}
func main() {

	conns := make([]net.Conn, 0)  // 用来存储所有的链接
	fmt.Println("服务器开始监听")
	// 使用tcp网络协议
	// 在本地监听8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("服务器监听打开失败，失败原因：", err)
		return
	}

	defer listen.Close()   // 延迟关闭listen

	// 循环等待客户端来连接
	for {
		// 等待客户端来链接我
		fmt.Println("等待客户端连接")
		conn, err := listen.Accept()    // 等待并返回连接对象，一直处于阻塞状态，直到有人来连接
		if err != nil {
			fmt.Println("客户端连接到服务器失败", err)
		} else {
			fmt.Println("Accept() 返回的连接对象 conn = ", conn)
			fmt.Println("连接客户的ip是：", conn.RemoteAddr().String())
			conns = append(conns, conn)
			fmt.Println(conns)
		}
		// 这里准备起一个协程，为本次连接过来的客户端服务
		go process(conn, conns)
	}
	fmt.Printf("listen = %v", listen)
}