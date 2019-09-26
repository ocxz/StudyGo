package main
import (
	"fmt"
	"os"
)

func main() {

	// 编写一段代码、获取命令行参数，打印参数个数，以及所有命令行参数

	fmt.Printf("命令行的参数有：%v个\n", len(os.Args))

	// 遍历os.Args切片，就能得到所有的命令行输入的参数值
	for i, v := range os.Args {
		fmt.Printf("args[%v] = %v\n", i, v)
	}
}