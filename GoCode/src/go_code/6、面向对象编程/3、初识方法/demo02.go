package main
import (
	"fmt"
)

type Circle struct {
	Radius float64
}

func (circle *Circle) Area() float64 {
	// 标准用法
	return 3.14 * (*circle).Radius * (*circle).Radius
	// 编译器在底层将circle.Radius转换成(*circle).Radius
	// return 3.14 * circle.Radius * circle.Radius
}

func main() {
	// 创建一个Circle变量
	c := Circle{4}
	// 标准用法
	res := (&c).Area()
	// 编译器在底层将c.Area()转换为（&c).Area()
	// res := c.Area()
	fmt.Printf("半径为%v的⚪的面积是%v\n", c.Radius, res)
}