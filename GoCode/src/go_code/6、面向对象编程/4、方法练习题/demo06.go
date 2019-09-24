package main
import (
	"fmt"
)

/*
	定义一个小小计数器结构体，实现加减乘除四个功能
	实现形式：分四个方法实现
	实现形式：分一个方法实现
*/

type Calcuator struct {
	Num1 float64
	Num2 float64
	Operator string
}

func (calcuator Calcuator)Calculate() float64 {
	switch calcuator.Operator{
	case "+":
		return calcuator.Num1 + calcuator.Num2
	case "-":
		return calcuator.Num1 - calcuator.Num2
	case "*":
		return calcuator.Num1 * calcuator.Num2
	case "/":
		return calcuator.Num1 / calcuator.Num2
	default:
		fmt.Println("运算符输入有误")
		return 0.0
	}
}

func main() {

	c := Calcuator{2, 5.4, "+"}
	fmt.Println(c.Calculate())
}