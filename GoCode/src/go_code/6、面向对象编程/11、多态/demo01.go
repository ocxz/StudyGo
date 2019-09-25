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
	Name string
}
// Phone实现Usb接口的方法
func (phone Phone) Start(){
	fmt.Println("手机开始工作")
}
func (phone Phone) Stop() {
	fmt.Println("手机停止工作")
}

type Camera struct {
	Name string
}
// 相机实现Usb接口的方法
func (camera Camera) Start(){
	fmt.Println("相机开始工作")
}
func (phone Camera) Stop() {
	fmt.Println("相机停止工作")
}

func main() {

	// 定义一个Usb接口数组，可以存放Phone和camear的结构体变量
	// 这里就可以体现多态数组
	usbs := [...]Usb{
		Phone{"vivo手机"},
		Camera{"尼康相机"},
		Phone{"小米手机"},
	}

	fmt.Println(usbs)
}