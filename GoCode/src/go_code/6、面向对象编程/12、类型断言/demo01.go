package main
import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {

	var a interface{} = Point{1, 2}

	// 类型断言，相当于as 强制类型转换，转换失败则会报错
	b := a.(Point)
	fmt.Println(b)
}