package main
import (
	"fmt"
)

type Student struct {
	Name string
	Age int
}
// 给Student实现String()方法
func (stu *Student) String() string {
	return fmt.Sprintf("Name = %v，Age = %v", stu.Name, stu.Age)
}

func main() {

	// 定义一个Student变量
	stu := Student{"张三", 18}
	fmt.Println(&stu)  // 传入指针，找到和结构体指针绑定的String()方法，返回执行
	fmt.Println(stu.String())
}