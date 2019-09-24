package main
import (
	"fmt"
)

// 方法的声明和调用
type Person struct {
	Name string
}
// 表示A结构体，有一个方法交test()
func (person Person) test() {
	fmt.Println("test()-->", person.Name)
}

func (person Person) speak() {
	fmt.Println(person.Name,"你是个好人~")
}

func (p Person) jisuan(num int){
	res := 0
	for i := 0; i <= num; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算的结果是", res)
}

func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}
func main() {

	person := Person{"张三"}
	person.test()
	person.speak()
	person.jisuan(100)
	res := person.getSum(10, 20)
	fmt.Println(res)
}