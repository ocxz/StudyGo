package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	// 打开已存在文件，将原来内容覆盖，写入新的内容

	path := "C:/Users/22527/Desktop/test2.txt"

	file, err := os.OpenFile(path, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("文件打开失败，失败原因：", err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	str := "你好，中国\r\n"
	for i :=0; i < 10; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
	fmt.Println("写入完成！")
}