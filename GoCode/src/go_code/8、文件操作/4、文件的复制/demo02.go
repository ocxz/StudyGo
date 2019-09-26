package main
import (
	"fmt"
	"io/ioutil"
)

func CopyFile(source string, target string) bool {

	// 将源文件读取到内存中   ioutil.ReadFile(path string)
	// 将内容写入到目标文件中  ioutil.WriteFile(filename string, data []byte, perm os.FileMode) error

	// data是一个[]byte
	data, err := ioutil.ReadFile(source)
	if err != nil {
		fmt.Println("源文件打开失败，失败原因：", err)
		return false
	}

	err = ioutil.WriteFile(target, data, 0666)
	if err != nil {
		fmt.Println("写入目标文件时出错，出错原因：", err)
		return false
	}

	return true
}

func main() {

	spath := "C:/Users/22527/Desktop/test2.txt"
	tpath := "C:/Users/22527/Desktop/test4.txt"
	if CopyFile(spath, tpath) {
		fmt.Println("文件复制成功")
	} else {
		fmt.Println("文件复制失败")
	}
}