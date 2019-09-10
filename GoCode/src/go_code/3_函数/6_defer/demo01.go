package main
import (
	"fmt"
)

func sum(n1 int, n2 int) int {

	// 当执行到defer时，暂时不执行，会将defer后面的语句压人到独立的栈中（可理解为defer栈）
	// 当函数执行完毕后，按照先进后出的方式出栈执行
	defer fmt.Println("ok n1=", n1)    // defer语句   3.打印ok1 n1=20
	defer fmt.Println("ok n2=", n2)     // 2.打印ok2 n2=30
	n1++
	n2++
	result := n1 + n2                   
	fmt.Println("ok result=", result)    // 1.打印ok result=50
	return result                  // 执行到这后，执行完毕，开始从defer栈出语句，先进后出执行
}

func main() {
	result := sum(20, 30)
	fmt.Println("main result=", result)
}