package main
import (
	"fmt"
)

// 要求：map的可以对应学生学号、map的值是一个结构体，包含姓名、年龄, 家庭住址

// 结构体声明
type Stu struct {
	Name string
	Age int
	Address string
}

func main() {

	// students := make(map[string]Stu)
	students := map[string]Stu{
		"stu3" : Stu{"张三", 18, "深圳"},
		"stu4" : Stu{"李四", 28, "广州"},
	}
	// 创建两个学生
	stu1 := Stu{ Name : "tom", Age : 18, Address : "北京",}
	stu2 := Stu{ Name : "mary", Age : 20, Address : "上海",}
	students["stu1"] = stu1
	students["stu2"] = stu2
	fmt.Println(students)

	// 变量各个学生信息
	for k, student := range students {
		fmt.Println("学生编号为：", k)
		fmt.Println("\t学生姓名：",student.Name)
		fmt.Println("\t学生年龄：",student.Age)
		fmt.Println("\t学生住址：",student.Address)
	}
}