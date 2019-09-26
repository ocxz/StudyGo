package main
import (
	"fmt"
	"go_code/7、简单项目开发/2、客户关系管理/service"
	"go_code/7、简单项目开发/2、客户关系管理/model"
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
	customers := this.cs.List()
	// 显示
	fmt.Println("----------------------------客户列表----------------------------\n")
	// 表头打印
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱\n")
	for _, cus := range customers {
		fmt.Println(cus.GetInfo())
	}
	fmt.Println("\n---------------------------客户列表完成---------------------------\n")
}

// 得到用户的输入信息，构建新的用户，并完成添加
func (this *customerView) add() {

	name := ""
	gender := ""
	age := 0
	phone := ""
	email := ""
	fmt.Println("----------------------------添加客户----------------------------\n")
	fmt.Print("姓名：")
	fmt.Scanln(&name)
	fmt.Print("性别：")
	fmt.Scanln(&gender)
	fmt.Print("年龄：")
	fmt.Scanln(&age)
	fmt.Print("手机号码：")
	fmt.Scanln(&phone)
	fmt.Print("邮箱：")
	fmt.Scanln(&email)

	// 构建一个新的Customer实例，Id号是唯一的，没有让用户输入，需要系统分配
	cus := model.NewCustomer2(name, gender, age, phone, email)
	// 调用Service中的add方法，添加客户
	if this.cs.Add(cus) {
		fmt.Println("\n---------------------------客户添加完成---------------------------\n")
	} else {
		fmt.Println("\n---------------------------客户添加失败---------------------------\n")
	}
}

// 得到用户输入，删除该id对应的客户
func (this *customerView) delete() {

	id := -1
	choice := ""
	fmt.Println("----------------------------删除客户----------------------------\n")
	fmt.Print("请输入要删除客户的编号（-1退出删除）：")
	fmt.Scanln(&id)
	if id == -1 {
		return   // 放弃删除
	}

	// 询问用户是否删除，循环询问
	for {
		fmt.Print("确定是否删除（y/n）：")
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" || choice == "Y" || choice == "N" {
			break
		}
	}
	// 调用service中的Delete方法
	if choice == "y" || choice == "Y" {
		if this.cs.Delete(id) {  // 删除成功
			fmt.Println("\n---------------------------客户删除完成---------------------------\n")
		} else {
			fmt.Println("\n-------------------客户删除失败，输入的id号不存在-------------------\n")
		}
	}
}

// 得到用户输入，完成修改功能
func (this *customerView) update() {

	id := -1
	fmt.Println("\n----------------------------修改客户----------------------------\n")
	fmt.Print("请选择要修改客户的编号（-1退出）：")
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	// 获取要修改的对象
	cus := this.cs.FindCustomerById(id)
	if cus.Id != 0 {
		this.key = ""
		fmt.Printf("姓名（%v）：", cus.Name) 
		fmt.Scanln(&this.key)
		if this.key != "" {
			cus.Name = this.key
		}

		this.key = ""
		fmt.Printf("性别（%v）：", cus.Gender) 
		fmt.Scanln(&this.key)
		if this.key != "" {
			cus.Gender = this.key
		}

		age := 0
		fmt.Printf("年龄（%v）：", cus.Age) 
		fmt.Scanln(&age)
		if age != 0 {
			cus.Age = age
		}

		this.key = ""
		fmt.Printf("电话（%v）：", cus.Phone) 
		fmt.Scanln(&this.key)
		if this.key != "" {
			cus.Phone = this.key
		}

		this.key = ""
		fmt.Printf("邮箱（%v）：", cus.Email) 
		fmt.Scanln(&this.key)
		if this.key != "" {
			cus.Email = this.key
		}

		// 更新
		if this.cs.Update(cus) {
			fmt.Println("\n---------------------------客户更新完成---------------------------\n")
		} else {
			fmt.Println("\n-------------------客户更新失败，输入的id号不存在-------------------\n")
		}
	} else {
		fmt.Println("\n-------------------客户更新失败，输入的id号不存在-------------------\n")
	}
}

// 退出系统，循环提问是否退出
func (this *customerView) exit() {
	
	for {
		fmt.Print("确定是否退出系统（y/n)：")
		fmt.Scanln(&this.key)
		if this.key == "y" || this.key == "Y" || this.key == "n" || this.key == "N" {
			break
		}
		fmt.Println("输入有误，请重新输入")
	}
	if this.key == "n" || this.key == "N" {
		fmt.Println("取消退出")
	} else {
		this.loop = false
	}
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
			this.add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
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