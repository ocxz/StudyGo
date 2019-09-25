package main
import (
	"fmt"
)

type Student struct {
	name string
	gender string
	age int
	id int
	score float64
}

func (stu *Student) say() string {

	info := fmt.Sprintf("student的信息 name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]", 
	stu.name, stu.gender, stu.age, stu.id, stu.score)
	return info
}

func main() {

	stu := Student{"张三", "男", 18, 1, 82.5}
	fmt.Println(stu.say())
}