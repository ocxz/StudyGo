package main
import (
	"fmt"
	"bufio"
	"os"
	"io"
)

	// 目标：统计英文、数字、空格、以及其它字符数量
	// 思路：打开文件，创建Reader
	// 每读一行，判断一次统计英文、数字、空格、以及其它字符数量
	// 将结果保存到结构体中

type CharCount struct {
	LetterCount int  // 记录英文字母的个数
	CNCount int   // 中文个数
	NumCount int     // 记录数字的个数
	SpaceCount int  // 记录空格的个数
	OtherCount int  // 记录其它字符的个数
}
func main() {

	path := "C:/Users/22527/Desktop/test.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("文件打开失败，失败原因", err)
		return
	}
	defer file.Close()

	// 定义一个CharCount实例
	var count CharCount
	// 创建一个Reader
	reader := bufio.NewReader(file)

	// 逐行循环读取内容
	for {
		content, err := reader.ReadString('\n')
		for _, v := range content {
			switch {
				case v >= 'a' && v < 'z':
					fallthrough  // 穿透，只穿一层
				case v >= 'A' && v < 'Z':
					count.LetterCount++
				case v == ' ' || v == '\t':
					count.SpaceCount++
				case v >= '0' && v <= '9':
					count.NumCount++
				case v >= 16640 && v <= 40869:
					count.CNCount++
				default:
					count.OtherCount++
			}
		}
		if err == io.EOF {
			break
		}
	}
	fmt.Println("英文字母的个数：", count.LetterCount)
	fmt.Println("汉字的个数：", count.CNCount)
	fmt.Println("数字的个数：", count.NumCount)
	fmt.Println("空格的个数：", count.SpaceCount)
	fmt.Println("其它字符的个数：", count.OtherCount)
}