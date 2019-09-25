package main
import (
	"fmt"
)

// 动物接口
type AnimalImpl interface {
	Eat()
}

// 能跑的接口
type RunableImp interface {
	Run()
}

// 人类的接口
type PersonImpl interface {
	AnimalImpl
	RunableImp
	Study()
}

type Student struct {
	Name string
}

// 实现AnimalImpl接口
func (stu Student) Eat() {
	fmt.Println(stu.Name + "：我能吃")
}

// 实现能跑的接口
func (stu Student) Run() {
	fmt.Println(stu.Name + "：我能跑")
}

// run() + eat() + study() 实现了人类的接口
func (stu Student) Study() {
	fmt.Println(stu.Name + "：我能学习")
}

func main() {

	var p PersonImpl = Student { "张三"}
	p.Eat()
	p.Run()
	p.Study()
}