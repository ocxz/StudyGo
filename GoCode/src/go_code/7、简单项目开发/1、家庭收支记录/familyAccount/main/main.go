package main
import (
	"fmt"
	"go_code/7、简单项目开发/1、家庭收支记录/familyAccount/utils"
)

func main() {

	fmt.Println("这个是面向对象的方式完成")
	utils.NewFamilyAccount().MainMenu()
}