package main
import (
	"fmt"
)

/*
	定义一个立方体，长宽高属性，定义一个计算体积的方法
*/

type Box struct {
	Length float64
	Width float64
	Height float64
}

func (box *Box) GetVolumn() float64 {
	return box.Length * box.Width * box.Height
}

func main() {

	// 定义一个长方体
	var box Box
	fmt.Printf("请输入长方体的长：")
	fmt.Scanln(&box.Length)
	fmt.Printf("请输入长方体的宽：")
	fmt.Scanln(&box.Width)
	fmt.Printf("请输入长方体的高：")
	fmt.Scanln(&box.Height)

	fmt.Printf("长方体%v*%v*%v的体积为%v", box.Length, box.Width, box.Height, box.GetVolumn())

}