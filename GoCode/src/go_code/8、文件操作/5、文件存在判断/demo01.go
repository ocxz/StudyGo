package main
import (
	"fmt"
	"os"
)

// 判断文件是否存在，如果返回err，则说明判断失败，无法确定文件是否存在
func PathIsExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {   // 文件确定已存在
		return true, nil
	}
	if os.IsNotExist(err) {   // 确定文件不存在
		return false, nil
	} 
	return false, err   // 其它错误，但不确定是否存在
}

func main() {

	path := "C:/Users/22527/Desktop/test20.txt"
	exist, err := PathIsExists(path)
	if err != nil {
		fmt.Println("判断失败，失败原因：", err)
	} else {
		if exist {
			fmt.Println("文件已存在")
		} else {
			fmt.Println("文件不存在")
		}
	}
}