package main
import (
	"fmt"
	"os"
	"io"
	"bufio"
)

// 自己写一个函数，接收两个文件路径 srcPath、dstPath
func CopyFile(srcPath string, dstPath string) (written int64, err error) {

	// 获取Reader
	sfile, err := os.Open(srcPath)
	if err != nil {
		fmt.Println("源文件打开失败，失败原因：", err)
		return 0, err
	}
	defer sfile.Close()

	// 通过bufio.NewReader(file *File) 获得Reader对象
	reader := bufio.NewReader(sfile)

	// 获取Reader
	tfile, err := os.OpenFile(dstPath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("目标文件打开失败，失败原因：", err)
		return 0, err
	}
	defer tfile.Close()

	// 通过bufio.NewWriter(file *File) 获得writer对象
	writer := bufio.NewWriter(tfile)

	// 真正的拷贝
	return io.Copy(writer, reader)
}
func main() {

	// 拷贝二进制文件
	srcpath := "F:/图片/IMG_20181117_212230.jpg"
	tpath := "C:/Users/22527/Desktop/my.jpg"
	_, err := CopyFile(srcpath, tpath)
	if err != nil {
		fmt.Println("文件拷贝失败")
	} else {
		fmt.Println("文件拷贝成功")
	}
}

