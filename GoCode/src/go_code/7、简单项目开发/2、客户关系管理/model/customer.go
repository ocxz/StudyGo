package model
import (
	"fmt"
)

// 声明一个Custormer结构体，表示客户信息
type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

func NewCustomer0() Customer {
	return Customer{}
}
// 使用工厂模式，返回一个Customer的实例
func NewCustomer(id int, name string, gender string,
	 age int, phone string, email string) *Customer{
	 return &Customer {
		 Id : id,
		 Name : name,
		 Gender : gender,
		 Age : age,
		 Phone : phone,
		 Email : email,
	 }
}

// 第二种创建Customer实例的方法
func NewCustomer2(name string, gender string,
	age int, phone string, email string) *Customer{
	return &Customer {
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
	}
}

// 写一个方法，显示该用户的信息，返回格式化的字符串
func (this *Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", this.Id, this.Name, this.Gender,
	 this.Age, this.Phone, this.Email)
	 return info
}