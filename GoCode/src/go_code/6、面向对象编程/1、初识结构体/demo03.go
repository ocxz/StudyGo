package main
import (
	"fmt"
)

// 创建结构体
type Person struct {
	Name string
	Age int
}

func main() {

	// 方式二
	p1 := Person{}
	p1.Name = "marry"
	p1.Age = 16
	p2 := Person{ Name : "tom", Age : 18 }

	fmt.Println(p1)
	fmt.Println(p2)

	// 方式三，使用指针的方式
	var p3 *Person = new(Person)  // p3 是一个指针
	// 标准写法，访问结构体变量
	(*p3).Name = "smit"
	(*p3).Age = 18
	fmt.Println(*p3)
	// 等价于下面的写法
	// go的设计者为了程序员使用方便，顶层会对p3.Name进行处理，给p4加上去做运算(*p4).Name = "张三"
	p4 := new(Person)  // 注意p4是一个指针
	p4.Name = "张三"
	p4.Age = 19
	fmt.Println(*p4)

	// 方式四，原理 = 方式二 + 方式三
	var person2 *Person = &Person{}
	person2.Name = "王五"
	person2.Age = 18
	fmt.Println(*person2)

	person3 := &Person{}
	(*person3).Name = "赵六"
	(*person3).Age = 21
	fmt.Println(*person3)

	person5 := &Person{"田七", 21}
	fmt.Println(*person5)
}