package main
import (
	"fmt"
	"os"
)

func main() {

	// 打开一个文件，这里的file是一个结构体指针，File
	// 1、file 交 file对象
	// 2、file 叫 file指针
	// 3、file 叫 file 文件句柄
	file, err := os.Open("C:/Users/22527/Desktop/test.txt")
	if err != nil {
		fmt.Println("open file err：", err)
	}

	// 输出文件，看看文件是什么
	fmt.Println("file=", *file)

	// 关闭
	err = file.Close()
	if err != nil {
		fmt.Println("close file err：", err)
	}
}