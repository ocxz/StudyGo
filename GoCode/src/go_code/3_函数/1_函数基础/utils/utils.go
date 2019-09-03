package utils   // 打包
import (
	"fmt"
)

// 将计算的功能代码块，放入一个函数中，当需要时调用即可
func Calculate(n1 float64, n2 float64, operator byte) float64 {

	switch operator {
	case '+':
		return n1 + n2
	case '-':
		return n1 - n2
	case '*':
		return n1 * n2
	case '/':
		return n1 / n2
	default:
		fmt.Println("操作符输入有误，无法计算")
	}
	return 0.0
}