package main

import (
	"fmt"
)

func main() {

	// 编写程序，可接受用户输入的年龄，大于18输出你的年龄大于18，要对自己的行为负责，否则输出你的年龄小于18，这次就放你一马
	var age byte
	fmt.Print("请输入你的年龄：")
	fmt.Scanln(&age)
	// fmt.Scanf("%d", &age)
	if age >= 18 {
		fmt.Println("你的年龄大于等于18，要对自己的行为负责")
	} else {
		fmt.Println("你的年龄小于18，这次就放你一马")
	}
}