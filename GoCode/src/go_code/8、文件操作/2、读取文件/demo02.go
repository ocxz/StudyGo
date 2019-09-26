package main
import (
	"fmt"
	"io/ioutil"
)

func main() {

	// 使用ioutil.ReadFile将文件一次性读取到位，ReadFile方法包含了文件打开关闭操作
	path := "C:/Users/22527/Desktop/test.txt"
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("文件读取失败", err)
	}

	// 将读取到的内容显示到终端
	fmt.Println(string(content))   // []byte
}