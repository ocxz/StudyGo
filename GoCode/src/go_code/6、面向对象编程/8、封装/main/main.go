package main
import (
	"fmt"
	"go_code/6、面向对象编程/8、封装/model"
)

func main() {

	acc := model.CreateAccount("123456", "123456", 58)
	fmt.Println(*acc)
}