package main
import (
	"fmt"
	"go_code/6、面向对象编程/9、继承/model"
)

func main() {

	// 使用嵌入了匿名结构体的结构体
	// 小学生结构体
	pupil := &model.Pupil{}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 8

	pupil.Testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()

	// 大学生结构体
	graduate := &model.Graduate{}
	graduate.Student.Name = "mary"
	graduate.Student.Age = 20

	graduate.Testing()
	graduate.Student.SetScore(90)
	graduate.Student.ShowInfo()

	book := &model.Book{}
	book.Goods.Name = "红楼梦"
	book.Goods.Price = 18
	book.Writer = "曹雪芹"
	fmt.Println(*book)

	// 结构体变量创建时，可直接指定各个匿名结构体字段的值
	tv := &model.TV{
		&model.Goods{"电视机", 5000},
		&model.Brand{"海尔", "山东"},
	}
	// 底层转换-->*((*tv).Goods)
	fmt.Println(*((*tv).Goods), *tv.Brand)
}