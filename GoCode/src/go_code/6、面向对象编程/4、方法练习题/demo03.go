package main
import (
	"fmt"
	"strconv"
)

/*
	要求：编写一个矩形结构体，指名长宽，为其绑定一个Area()方法，返回面积，main函数中调用，并打印
*/

type Rect struct {
	Width float64
	Height float64
}

func (rect *Rect) Area() float64 {
	res := rect.Width * rect.Height
	r, _ := strconv.ParseFloat(fmt.Sprintf("%.2f",res), 64)
	return r
}

func main() {

	rect := Rect{3.14, 4.56}
	fmt.Println(rect.Area())
}