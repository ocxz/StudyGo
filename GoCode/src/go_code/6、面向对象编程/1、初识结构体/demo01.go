package main
import (
	"fmt"
)

// 定义一个Cat结构体，将Cat的各个字段放入Cat结构体进行管理
type Cat struct {
	Name string  // 默认""
	Age int      // 默认0
	Color string
	Scores []int
}
func main() {

	// 创建两只猫
	var cat1 Cat   // 自定义类型，类似于var a int
	cat1.Name = "小白"
	cat1.Age = 18
	cat1.Color = "白色"
	fmt.Println("cat1=", cat1)
}