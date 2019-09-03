package main

import (
	"fmt"
)

func main() {

	// 演示goto的使用
	n := 10
	fmt.Println("ok1")
	fmt.Println("ok2")
	// lable1:
	if n >5 {
		goto label3
	}
	fmt.Println("ok3")
	fmt.Println("ok4")
	// label2:
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")
	label3:
	fmt.Println("ok8")
	fmt.Println("ok9")
}