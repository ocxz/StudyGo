package main
import (
	"fmt"
	"net"   // 做网络socket开发时，net包含我们所需的所有方法和函数
	_ "io"
)

// 处理客户端的请求
func process(conn net.Conn) {

	// 这里我们循环接收客户端发送的数据
	defer conn.Close()   // 关闭conn连接
	fmt.Printf("服务器在等待客户端%v的输入\n", conn.RemoteAddr().String())
	for {
		// 创建一个新的切片
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)    // 客户端没有发送数据，就会一直阻塞，如果客户端没有write，那么协程就会阻塞在这里
		// if err == io.EOF {
		// 	fmt.Println("客户端退出")
		// 	return
		// }
		if err != nil {
			fmt.Println("服务端读取客户端发送的数据失败，失败原因：", err)
		}
		msg := string(buf[:n])   // 客户端发过来的信息
		if msg == "exit" {
			fmt.Printf("客户端%v结束了会话\n", conn.RemoteAddr().String())
			return
		}
		// 显示客户端发送的数据到服务器终端
		// 陷阱：1、不需要println，因为客户端发送过来的一行中，已经有\n换行符了
		//      2、只能使用切片来接收，如果要显示需要将其转为string
		//      3、切片必须是 :n，防止将切片后面乱起八糟的数据，也给转换输出了
		fmt.Print(string(buf[:n]))
	}
}
func main() {

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
		}
		// 这里准备起一个协程，为本次连接过来的客户端服务
		go process(conn)
	}
	fmt.Printf("listen = %v", listen)
}