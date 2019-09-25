package main
import (
	"fmt"
	"go_code/7、简单项目开发/2、客户关系管理/service"
)

type customerView struct {

	// 定义必要的字段
	key string  // 接收用户输入
	loop bool  // 表示是否循环显示菜单
	// 增加一个字段customerService
	cs *service.CustomerService
}

// 显示所有的客户信息
func (this *customerView) list() {

	// 首先获取当前所有客户信息，在切片中
	customers := this.CustomerService.List()
	// 显示
	fmt.Println("----------------------------客户列表----------------------------")
	// 表头打印
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _, cus := range customers {
		fmt.Println(cus.GetInfo())
	}
	fmt.Println("---------------------------客户列表完成---------------------------")
}

// 显示主菜单
func (this *customerView) mainMenu() {

	for {
		fmt.Println("\n--------------------------客户信息管理软件--------------------------\n")
		fmt.Println("                       1 添 加 客 户")
		fmt.Println("                       2 修 改 客 户")
		fmt.Println("                       3 删 除 客 户")
		fmt.Println("                       4 客 户 列 表")
		fmt.Println("                       5 退 出 系 统")
		fmt.Println()
		fmt.Print("      请选择（1-5）：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			fmt.Println("添 加 客 户")
		case "2":
			fmt.Println("修 改 客 户")
		case "3":
			fmt.Println("删 除 客 户")
		case "4":
			fmt.Println("客 户 列 表")
		case "5":
			this.loop = false
		default:
			fmt.Println("输入有误，请重新输入")
		}
		if !this.loop {   // 跳出循环
			break
		}
	}
	fmt.Println("退出客户关系管理系统")
}

func main() {

	// 创建CustomerView，并运行显示主菜单
	cv := customerView {
		key : "",
		loop : true,
		cs : service.NewCustomerService(),
	}
	// 显示主菜单
	cv.mainMenu()
}