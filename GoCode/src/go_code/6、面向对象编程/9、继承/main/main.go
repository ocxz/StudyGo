package main
import (
	"fmt"
	"go_code/6、面向对象编程/9、继承/model"
)

func main() {

	stu := model.Pupil{ "张三", 18, 59}
	stu.Testing()
	stu.SetScore(95)
	stu.ShowInfo()
	fmt.Println(stu)
}