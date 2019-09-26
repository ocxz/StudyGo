package service
import (
	"go_code/7、简单项目开发/2、客户关系管理/model"
)

// 该CustomerService完成对Customer的操作，完成增删改查
type CustomerService struct {

	// 用来存储所用的customer用户
	customers []*model.Customer
	// 声明一个字段，表示当前切片含义多少个客户
	// 后可作为新客户的编号：customerNum+1
	customerNum int
}

// 编写一个方法，可以返回 CustomerService
// 初始化
func NewCustomerService() *CustomerService {

	// 为了能够，看到有客户在切片中，我们初始化一个客户
	cs := CustomerService{
		customerNum : 1,
		customers : []*model.Customer{
			model.NewCustomer(1, "张三", "男", 20, "18370810801", "zs@sohu.com"),
		},
	}
	return &cs
}

// 编写List方法，返回客户切片
func (this *CustomerService) List() []*model.Customer{
	return this.customers
}

// 添加客户到[]*model.Customer中
// 一定要用指针绑定，要不然值拷贝，不会操作到原来的CustomerService
func (this *CustomerService) Add(customer *model.Customer) bool {

	// 我们确定一个分配id的规则，就是添加顺序
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

// 根据id查找客户在切片中对应下标，如果没有该客户，返回-1
func (this *CustomerService) FindById(id int) int {

	// 遍历切片，看看有没有对应Id的客户
	index := -1
	for i :=0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {   // 找到了
			index = i
		}
	}
	return index
}

// 根据id查找客户，返回一个客户，没有找到返回一个初始化了的Customer
func (this *CustomerService) FindCustomerById(id int) model.Customer {

	index := this.FindById(id)
	if index == -1 {   // 没有找到，返回初始化了的Customer
		return model.NewCustomer0()
	} 
	return *this.customers[index]
}

// 根据id删除客户（从切片中删除）
func (this *CustomerService) Delete(id int) bool {

	// 找到要删除客户在切片中的下标
	index := this.FindById(id)
	if index == -1 {     // 客户不存在
		return false
	} else {

		// 然后从切片中删除元素
		this.customers = append(this.customers[:index], this.customers[index+1:]...)
		return true
	}
}

// 根据id来修改客户（在切片中修改即可）
func (this *CustomerService) Update(cus model.Customer) bool {

	// 找到要修改的客户下标
	index := this.FindById(cus.Id)
	if index == -1 {   // 要修改的客户不存在，返回false
		return false
	}
	this.customers[index] = &cus
	return true
}