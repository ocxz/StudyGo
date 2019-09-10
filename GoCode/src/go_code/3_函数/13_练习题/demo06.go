package main
import (
	"fmt"
	"errors"
)

// 两个操作数，一个操作符，完成计算器操作
// 1 加 2减 3乘 4除 0退出
func operate(num1 float64, num2 float64, code int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	switch code {
	case 1:
		fmt.Printf("%v + %v = %v \n", num1, num2, num1 + num2)
	case 2:
		fmt.Printf("%v - %v = %v \n", num1, num2, num1 - num2)
	case 3:
		fmt.Printf("%v * %v = %v \n", num1, num2, num1 * num2)
	case 4:
		if num2 == 0 {
			err := errors.New("操作有误，0不能作为除数")
			panic(err)
		}
		fmt.Printf("%v / %v = %v \n", num1, num2, num1 / num2)
	default:
		err := errors.New("操作码不正确")
		panic(err)
	}
}

func main() {
	println("---------------小小计算器---------------\n")
	println("1. 加法")
	println("2. 减法")
	println("3. 乘法")
	println("4. 除法")
	println("0. 退出")
	var num1, num2 float64   // 两个操作数
	var code int        // 操作码
	for {
		// fmt.Print("请输入两个操作数（空格隔开）：")
		// fmt.Scanf("%f %f", &num1, &num2)
		fmt.Print("请输入第一个操作：")
		fmt.Scanln(&num1)
		fmt.Print("请输入第二个操作：")
		fmt.Scanln(&num2)
		fmt.Print("请输入操作码：")
		fmt.Scanln(&code)
		if code == 0 {
			fmt.Println("程序退出！")
			break
		} else {
			operate(num1, num2, code)
		}
	}
}