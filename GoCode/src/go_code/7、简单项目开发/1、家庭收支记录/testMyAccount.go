package main
import (
	"fmt"
)

func main() {

	// 定义账户余额
	balance := 10000.0
	// 每次收支的金额
	money := 0.0
	// 每次收支的说明
	note := ""
	// 收支详情
	// 当有收支时，只需要对details进行拼接处理
	details := "收支\t账户金额\t收支金额\t说    明"

	// 定义一个变量，标明是否有收支的行为
	flag := false

	// 声明一个变量，保存接收用户输入选项
	key := ""
	// 声明变量，控制循环
	loop := true
	// 显示这个主菜单，循环显示
	for {
		fmt.Println("--------------------------家庭收支记账软件--------------------------")
		fmt.Println()
		fmt.Println("                            1  收支明细")
		fmt.Println("                            2  登记收入")
		fmt.Println("                            3  登记支出")
		fmt.Println("                            4  退    出")
		fmt.Println()
		fmt.Print("                            请选择（1-4）")
		fmt.Scanln(&key)
		switch key {
			case "1":
				fmt.Println("\n--------------------------当前收支明细记录--------------------------\n")
				if flag {
					fmt.Println(details)
				} else {
					fmt.Println("当前没有收支明细")
				}
			case "2":
				fmt.Print("本次收入金额：")
				fmt.Scanln(&money)
				balance += money   // 修改账户余额
				fmt.Print("本次收入说明：")
				fmt.Scanln(&note)
				// 将收入情况，拼接到details变量中
				flag = true
				details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance, money, note)

			case "3":
				fmt.Print("本次支出金额：")
				fmt.Scanln(&money)
				// 这里需要做出money是否大于balance，否则透支
				if money > balance {
					fmt.Println("余额不足")
					break
				}
				balance -= money
				fmt.Print("本次支出说明：")
				fmt.Scanln(&note)
				flag = true
				details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", balance, money, note)
			case "4":
				choice := ""
				fmt.Print("你确定要退出吗？y/n：")
				fmt.Scanln(&choice)
				for {
					if choice == "y" {
						loop = false
						break
					} else if choice == "n"{
						break
					} else {
						fmt.Println("输入有误，请重新输入")
						fmt.Print("你确定要退出吗？y/n：")
						fmt.Scanln(&choice)
					}
				}
			default:
				fmt.Println("请输入正确的选项")
		}

		if !loop {   // 退出for循环
			break
		}
	}
	fmt.Println("退出家庭记账软件的使用")

}