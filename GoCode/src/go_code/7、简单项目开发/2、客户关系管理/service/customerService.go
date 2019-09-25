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
func NewCustomerService() *CustomerService {

	// 为了能够，看到有客户在切片中，我们初始化一个客户
	cs：= &CustomerService {
		customerNum : 1,
		customers := []*model.Customer{
			NewCustomer(1, "张三", "男", 20, "18370810801", "zs@sohu.com"),
		}
	}
}

// 编写List方法，返回客户切片
func (this *CustomerService) List() []*model.Customer{
	return this.customers
}