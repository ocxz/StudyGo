package main
import (
	"fmt"
)

// 声明定义一个接口
type Usb interface {
	// 声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {

}
// Phone实现Usb接口的方法
func (phone Phone) Start(){
	fmt.Println("手机开始工作")
}
func (phone Phone) Stop() {
	fmt.Println("手机停止工作")
}

type Camera struct {

}
// 相机实现Usb接口的方法
func (camera Camera) Start(){
	fmt.Println("相机开始工作")
}
func (phone Camera) Stop() {
	fmt.Println("相机停止工作")
}

// 计算机
type Computer struct {

}
// 编写一个Working方法，接收接口类型数据
// 只要是实现了Usb接口，就是指实现了Usb接口声明的所有方法
func (computer *Computer)Working(usb Usb) {

	// 通过usb接口变量，来调用Start和Stop方法
	usb.Start()
	usb.Stop()
}

func main() {

	// 测试
	// 创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	// 关键点，调用Computer的working方法，传入实现接口Usb的结构体，调用相应的方法
	computer.Working(phone)
	computer.Working(camera)
}