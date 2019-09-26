package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	// 创建一个新文件，写入内容，带缓冲
	path := "C:/Users/22527/Desktop/test2.txt"
	file, err := os.OpenFile(path, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败，失败原因：", err)
		return
	}

	defer file.Close()   // 及时关闭

	// 准备写入5句话
	str := "hello，中国\r\n"
	// 写入时，使用带缓冲的写入对象 *Writer
	writer := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	// 因为writer是带缓冲的，因此调用WriterString方法时，是先写入缓冲[内存容器]
	// 所以需要Flush()，将缓冲中最后的内容Flush中写入存储中
	// 当然了，缓冲满了也会自动将缓冲区内容写入到文件中
	writer.Flush()

	fmt.Println("写入完成")
}