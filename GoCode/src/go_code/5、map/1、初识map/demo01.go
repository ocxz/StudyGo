package main
import (
	"fmt"
)

func main() {

	// map的声明和注意事项
	var m map[string]string
	// 使用map前需要先分配内存，如make初始化
	// make(数据类型，空间[数据个数])
	m = make(map[string]string, 10)
	m["No.1"] = "宋江"
	m["No.2"] = "卢俊义"
	fmt.Println(m)
}