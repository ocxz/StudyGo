package main
import (
	"fmt"
)

type Person interface {
	
	Say()
}

type Student struct {
	Name string
	Age int
}
func (stu Student) Say() {
	fmt.Printf("%v说你好，我今年%v岁了", stu.Name, stu.Age)
}
func (stu Student) Play() {
	fmt.Println(stu.Name + "正在玩游戏")
}

func main() {

	var p Person = Student{"张三", 18}
	p.Say()  // 可以调出来
	p.Play()  // 调不出来，因为person接口中没有定义Paly()方法
}