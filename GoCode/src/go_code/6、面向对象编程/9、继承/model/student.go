package model
import (
	"fmt"
)

// 共有属性
type Student struct {
	Name string
	Age int
	Score int
}

// 将共有的方法，绑定调共有匿名结构体中
func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

// 小学生
type Pupil struct {
	Student   // 嵌入Student结构体
}

// Pupil特有的方法，保留
func (p *Pupil) Testing() {
	fmt.Println("小学生正在考试")
}

type Graduate struct {
	Student  // 嵌入Student结构体
}

// Graduate特有的方法，保留
func (grad *Graduate) Testing() {
	fmt.Println("大学生正在考试")
}