package main

import (
	"fmt"
)

func main() {

	// break 小练习
	// 1、100以内数求和，求出 当和 第一次大于20的当前数

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println("和超过了20，当前数为：", i)
			break
		}
	}

	// 实现登录验证 有三次机会 如果用户为张无忌 密码为888 提示登陆成功，否则提示登陆失败
	time := 3
	
	for {
		var name string
		var password string
		fmt.Print("请输入用户名：")
		fmt.Scanln(&name)
		fmt.Print("请输入密码：")
		fmt.Scanln(&password)

		if name == "张无忌" && password == "888" {
			fmt.Println("登陆成功")
			break
		} else {
			time--
			if time <= 0 {
				fmt.Println("登陆的三次机会用完了，登陆失败")
				break
			} else {
				fmt.Printf("登陆失败，还有%d次登陆机会 \n", time)
			}
		}
	}
}