package model
import (
	"fmt"
)

type account struct {
	id string
	pwd string
	money float64
}

func CreateAccount(id string, pwd string, money float64) *account {

	if len(id) < 6 || len(id) > 10 {
		fmt.Println("账户格式不正确")
		return nil
	}

	if len(pwd) != 6  {
		fmt.Println("密码格式不正确")
		return nil
	} 

	if money <20  {
		fmt.Println("余额小于20")
		return nil
	}

	return &account{
		id : id,
		pwd : pwd,
		money : money,
	}
}

func (acc *account)SetId(id string) {
	if len(id) < 6 || len(id) > 10 {
		fmt.Println("账户格式不正确")
	} else {
		acc.id = id
	}
}
func (acc *account) GetId() string {
	return acc.id
}

func (acc *account)SetPwd(pwd string) {
	if len(pwd) != 6  {
		fmt.Println("密码格式不正确")
	} else {
		acc.pwd = pwd
	}
}
func (acc *account) GetPwd() string {
	return acc.pwd
}

func (acc *account)SetMoney(money float64) {
	if money <20  {
		fmt.Println("余额小于20")
	} else {
		acc.money = money
	}
}
func (acc *account) GetMoney() float64 {
	return acc.money
}
