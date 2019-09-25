package utils
import (
	"fmt"
)

type FamilyAccount struct {
		// 定义账户余额
		balance float64
		// 每次收支的金额
		money float64
		// 每次收支的说明
		note string
		// 收支详情
		// 当有收支时，只需要对details进行拼接处理
		details string
	
		// 定义一个变量，标明是否有收支的行为
		flag bool
	
		// 声明一个变量，保存接收用户输入选项
		key string
		// 声明变量，控制循环
		loop bool
		// 显示这个主菜单，循环显示
}

// 编写一个工厂模式的构造方法、返回一个*FamilyAccount
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key : "",
		loop : true,
		balance : 10000.0,
		money : 0.0,
		note : "",
		flag : false,
		details : "收支\t账户金额\t收支金额\t说    明",
	}
}

// 将显示明细写出一个方法
func (this *FamilyAccount) showDetails(){
	fmt.Println("\n--------------------------当前收支明细记录--------------------------\n")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细")
	}
}

// 将登记收入写成一个方法，和*FamilyAccount绑定
func (this *FamilyAccount) income() {
	fmt.Print("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money   // 修改账户余额
	fmt.Print("本次收入说明：")
	fmt.Scanln(&this.note)
	// 将收入情况，拼接到details变量中
	this.flag = true
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
}

// 将登记支出写成一个方法，和*FamilyAccount绑定
func (this *FamilyAccount) pay() {
	fmt.Print("本次支出金额：")
	fmt.Scanln(&this.money)
	// 这里需要做出money是否大于balance，否则透支
	if this.money > this.balance {
		fmt.Println("余额不足")
	} else {
		this.balance -= this.money
		fmt.Print("本次支出说明：")
		fmt.Scanln(&this.note)
		this.flag = true
		this.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	}
}

// 将退出系统写成一个方法，和*FamilyAccount绑定
func (this *FamilyAccount) exit() {
		choice := ""
		fmt.Print("你确定要退出吗？y/n：")
		fmt.Scanln(&choice)
		for {
			if choice == "y" {
				this.loop = false
				break
			} else if choice == "n"{
				break
			} else {
				fmt.Println("输入有误，请重新输入")
				fmt.Print("你确定要退出吗？y/n：")
				fmt.Scanln(&choice)
			}
		}
}

// 给该结构体绑定相应的功能
// 显示主菜单
func (this *FamilyAccount) MainMenu() {
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
		fmt.Scanln(&this.key)
		switch this.key {
			case "1":
				this.showDetails()
			case "2":
				this.income()
			case "3":
				this.pay()
			case "4":
				this.exit()
			default:
				fmt.Println("请输入正确的选项")
		}

		if !this.loop {   // 退出for循环
			break
		}
	}
	fmt.Println("退出家庭记账软件的使用")
}