package main
import (
	"fmt"
)

// 定义一个结构体
type Account struct {
	AccountNo string
	Pwd string
	Balance float64  // 余额
}
// 方法
// 1.可存款
func (account *Account) Deposite(money float64, pwd string){
	if account.Pwd == pwd {
		if money <= 0 {
			fmt.Println("输入的金额不正确")
		} else {
			account.Balance += money
			fmt.Println("存款成功")
		}
	}else {
		fmt.Println("输入的密码不正确")
	}
}
// 2.取款
func (account *Account) WithDraw(money float64, pwd string){
	if pwd != account.Pwd {
		fmt.Println("输入的密码不正确")
		return
	}

	if money <= 0 || money > account.Balance {
		fmt.Println("输入的金额不正确")
		return 
	}

	account.Balance -= money
	fmt.Println("取款成功")
}

// 3.查询余额
func (account *Account) QueryMoney(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("输入的密码不正确")
		return
	}

	fmt.Printf("账号%v余额为%v元\n", account.AccountNo, account.Balance)
}

func main() {

	// 定义一个账户变量
	account := &Account{
		AccountNo : "gs111111",
		Pwd : "123456",
		Balance : 10000,
	}
	fmt.Println(*account)
	account.QueryMoney("123456")
	account.Deposite(201.3, "123456")
	account.QueryMoney("123456")
	account.WithDraw(123.1, "123456")
	account.QueryMoney("123456")
}