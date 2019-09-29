package main
import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {

	// 建立与服务端的连接，服务器的ip是192.168.0.195
	ip := "192.168.0.195"
	conn, err := net.Dial("tcp", ip + ":8888")
	if err != nil {
		fmt.Printf("连接服务器%v失败，失败原因：%v\n", ip, err)
		return
	}
	fmt.Printf("连接服务器%v成功，连接对象是：%v\n", ip, conn)

	// 获取键盘标准输入流，即终端输入
	reader := bufio.NewReader(os.Stdin)       // os.Stdin执行标准输入流，即终端 /dev/stdin 其实是文件

	sum := 0
	for {
		// 从终端读取一行用户输入，并发送给服务器
		fmt.Print("我说：")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取终端输入失败，失败原因：", err)
			return
		}

		// 再将line方法送给服务器，即向连接中写入数据
		wn, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("向连接写入失败，失败原因：", err)
		}
		sum += wn
		if line == "exit\r\n" {
			fmt.Printf("结束向服务端发送数据，一共向服务器发送了%v个字节的数据\n", sum)
			break
		}

		// 接受服务端的响应，读取服务的相应的数据
		buf := make([]byte, 1024)
		rn, err := conn.Read(buf)
		if err != nil {
			fmt.Println("接受服务端响应信息失败，失败原因：", err)
		} else {
			fmt.Println("服务器说：", string(buf[:rn]))
		}

	}
}