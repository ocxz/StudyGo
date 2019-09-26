package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main() {

	/*

	*/

	file, err := os.Open("C:/Users/22527/Desktop/test.txt")
	if err != nil {
		fmt.Println("文件打开失败，失败原因：", err)
	}
	// 函数退出时，要及时关闭file
	defer file.Close() // 要及时关闭file句柄，否则会有内存泄漏

	// 创建一个 *Reader 是带缓冲的
	// 默认缓冲区大小是：4096个字节 defaultBufSize = 4096
	reader := bufio.NewReader(file)

	// 循环读取文件内容
	for {
		str, err := reader.ReadString('\n')   // 读到一个换行就结束一次

		// 输出内容
		fmt.Print(str)
		if err == io.EOF {   // io.EOF表示文件的末尾
			break
		}
	}
	fmt.Println("\n文件读取结束")
}
