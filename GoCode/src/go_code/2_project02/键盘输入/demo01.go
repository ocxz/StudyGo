package main

import (
	"fmt"
)

func main() {

	// 可从控制台接受用户信息  包括：姓名、年龄、薪水、是否通过
	// 两者方式：fmt.Scanln()、fmt.Scanf()获取

	// fmt.Scanln()
	var name string
	var age byte
	var sal float32
	var isPass bool
	// fmt.Println("请输入姓名")
	// // 当程序执行到fmt.Scanln()时，程序会停止，等待用户输入
	// fmt.Scanln(&name)

	// fmt.Println("请输入年龄")
	// fmt.Scanln(&age)

	// fmt.Println("请输入薪水")
	// fmt.Scanln(&sal)

	// fmt.Println("请输入是否通过考试")
	// fmt.Scanln(&isPass)

	// fmt.Println("名字是 ", name)
	// fmt.Println("年龄是 ", age)
	// fmt.Println("薪水是 ", sal)
	// fmt.Println("是否通过考试 ", isPass)


	// 方式二：fmt.Scanf()  可以按指定的格式输入
	fmt.Println("请输入你的姓名、年龄、薪水、是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Println("名字是 ", name)
	fmt.Println("年龄是 ", age)
	fmt.Println("薪水是 ", sal)
	fmt.Println("是否通过考试 ", isPass)


}