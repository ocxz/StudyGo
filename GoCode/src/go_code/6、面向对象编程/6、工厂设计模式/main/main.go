package main
import (
	"fmt"
	"go_code/6、面向对象编程/6、工厂设计模式/model"
)

func main() {

	// 如果是model.Student，公开的，则可以直接创建结构体变量
	// stu := model.Student{"张三", 18, 87.5}
	// fmt.Println(stu)

	// 也可调用函数，传入参数，返回结构体指针
	// 通过工厂模式来解决
	stu := model.CreateStudent("张三", 18, 20)
	   fmt.Println(*stu)
	stu.SetScore(36)
	   fmt.Println(stu.GetScore())
	   fmt.Println()
}