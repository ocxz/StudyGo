package main

import (
	"fmt"
)

func main() {

	// 出票系统
	// 4—10 旺季 成人(18-60)：60   儿童(<18)：半价   老人(>60)：1/3
	// 淡季 成人：40  其他：20

	// 两个条件，淡旺季、票类别

	var month byte
	var age byte
	var price float64
	fmt.Print("请输入月份：")
	fmt.Scanln(&month)
	fmt.Print("请输入年龄：")
	fmt.Scanln(&age)

	if month >= 4 && month <= 10 {   // 旺季

		if age >=18 && age <=60 {   // 旺季 成人
			fmt.Printf("您购买的是成人票，总价格为：%v元", price)
		} else if age < 18 {   // 旺季 儿童
			fmt.Printf("您购买的是儿童票，总价格为：%v元", price / 2)
		} else {
			fmt.Printf("您购买的是老年票，总价格为：%v元", price / 3)
		}
	} else {
		if age >=18 && age <=60 {   // 旺季 成人
			fmt.Printf("您购买的是成人票，总价格为：%v元", 40.00)
		} else {
			fmt.Printf("您购买的是其他类型的票，总价格为：%v元", 20.00)
		}
	}
}