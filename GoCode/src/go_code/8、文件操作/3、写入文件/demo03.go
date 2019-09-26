package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	// 打开已存在的文件，给文件内容追加，我爱中国
	path := "C:/Users/22527/Desktop/test2.txt"
	file, err := os.OpenFile(path, os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败，失败原因：", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	str := "我爱中国\r\n"
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	writer.Flush()

	fmt.Println("追加文件完成！")
}