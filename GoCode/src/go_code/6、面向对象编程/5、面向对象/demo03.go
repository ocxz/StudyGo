package main
import (
	"fmt"
)

/*
	题目：景区打印门票 姓名、年龄、门票价格打印
*/

type Visitor struct {
	Name string
	Age int
	Price float64
}

func (visitor *Visitor) PrintInfo() string {
	
	switch {
		case visitor.Age > 18:
			visitor.Price = 20
			return fmt.Sprintf("%v的年龄为：%v，门票价格为：%v", visitor.Name, visitor.Age, 20)
		default:
			visitor.Price = 0
			return fmt.Sprintf("%v的年龄为：%v，门票免费", visitor.Name, visitor.Age)
	}
}

func main() {

	var v Visitor
	for {
		fmt.Print("请输入姓名：")
		fmt.Scanln(&v.Name)
		if(v.Name == "n"){
			fmt.Println("退出程序")
			break
		}
		fmt.Print("请输入年龄：")
		fmt.Scanln(&v.Age)
		fmt.Println(v.PrintInfo())
	}
}