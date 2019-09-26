package main
import (
	"fmt"
	"bufio"
	"os"
	"io"
)

// 文件复制，使用缓冲区的方式
func CopyFile(soruce string, target string) bool {

	/*
		代码说明：
				1、打开源文件（defer 延迟关闭）   源文件 只读：os.O_RDONLY
				2、打开目标文件（defer 延迟关闭） 目标文件 只写+创建：os.O_WRONLY | os.O_CREATE
				3、创建读对象 reader := bufio.NewReader(file *File)
				4、创建写对象 writer := bufio.NewWriter(file *File)
				5、循环逐行读取  content, err := reader.ReadString("\n")
				6、在读取过程中，写入缓冲区 writer.WriteString(content)
				7、判断是否读取完毕  err == os.EOF
				8、将缓冲区剩余内容写入到文件中，writer.Flush()
	*/
	sfile, err := os.OpenFile(soruce, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("源文件打开失败，失败原因：", err)
		return false
	}
	defer sfile.Close()

	tfile, err := os.OpenFile(target, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("目标文件打开失败，失败原因：", err)
		return false
	}
	defer tfile.Close()

	// 创建读对象
	sReader := bufio.NewReader(sfile)
	// 创建写对象
	tWriter := bufio.NewWriter(tfile)

	// 循环逐行读取
	for {
		content, err := sReader.ReadString('\n')
		// 写入
		tWriter.WriteString(content)
		if err == io.EOF {
			break
		}
	}
	tWriter.Flush()
	return true
}
func main() {

	spath := "C:/Users/22527/Desktop/test2.txt"
	tpath := "C:/Users/22527/Desktop/test3.txt"
	if CopyFile(spath, tpath) {
		fmt.Println("文件复制成功")
	} else {
		fmt.Println("文件复制失败")
	}
}