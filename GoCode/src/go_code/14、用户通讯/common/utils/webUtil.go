package utils
import (
	"net"
	"io"
	"errors"
	"encoding/json"
	"encoding/binary"
	"go_code/14、用户通讯/common/message"
	"fmt"
)

// 传输者，负责消息的写入和读出
// 要有两个字段，conn、buf
type transfer struct {

	Conn net.Conn
	Buf [8096]byte
}
// 创建transfer传输者
func CreateTransfer(conn net.Conn) *transfer {
	return &transfer{ Conn : conn }
}

// 这是一个读包函数，给连接，读取长度和包内容
func (this *transfer) ReadPkg() (msg message.Message, err error) {

	// buf := make([]byte, 8096)
	fmt.Println("等待客户端发送数据")

	// 读取发过来的数据长度
	// conn只有在都没有关闭的情况下才会阻塞，任意一方关闭，都不会在阻塞
	n, err := this.Conn.Read(this.Buf[:4])
	if n != 4 || err != nil {
		if err == io.EOF {   // 对方关闭了连接
			return
		}
		err = errors.New(fmt.Sprintf("读取长度失败%v，失败原因：%v", n, err))
		return
	}
	// 将数据长度转换为uint32类型数据
	pkgLen := binary.BigEndian.Uint32(this.Buf[:4])

	// 读取消息内容本身
	n, err = this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New(fmt.Sprintf("消息内容读取失败，失败原因：%v", err))
		return
	}
	// 将读取到的内容反序列化成msg
	err = json.Unmarshal(this.Buf[:pkgLen], &msg)
	if err != nil {
		err = errors.New(fmt.Sprintf("内容反序列化失败，失败原因：%v", err))
		return
	}
	return
}
// 这是一个写包函数，向连接中写入包数据
func  (this *transfer) WritePkg(data []byte) (err error) {

	// 先发送长度给对方
	// 1、先把data的长度发送给对方
	pkgLen := uint32(len(data))
	// 2、将pkgLen转成字节切片
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)
	// 发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		err = errors.New(fmt.Sprintf("发送数据长度失败，失败原因：%v", err))
		return
	}

	// 发送信息结构体本身
	// 发送消息本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		err = errors.New(fmt.Sprintf("发送数据失败，失败原因：%v", err))
		return
	}
	return

}