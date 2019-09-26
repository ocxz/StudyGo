package main
import (
	"fmt"
	"bufio"
	"os"
	"io"
)

func main() {

	// 打开已存在的文件，将原来内容读出，显示在终端，且追加5句我爱北京
	path := "C:/Users/22527/Desktop/test2.txt"
	file, err := os.OpenFile(path, os.O_APPEND | os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("文件打开失败，失败原因：", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		if err == io.EOF {  // 读取到文件末尾
			break
		}
	}
	
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString("我爱北京\r\n")
	}
	writer.Flush()

	fmt.Println("追加完成")
}