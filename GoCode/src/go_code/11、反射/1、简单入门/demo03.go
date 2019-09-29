package main
import (
	"reflect"
	"fmt"
)

func reflectTest(b interface{}) {

	// 获取变量的rValue
	rValue := reflect.ValueOf(b)
	// 通过rValue.SetInt(20)，改变传进来的int类型的值

	// rValue.Elem(v)返回接口或者指针真实指向的值
	// 如：
	// rValue.SetInt(20)
	rElem := rValue.Elem()
	fmt.Printf("rElem = %v， rElem的类型是：%T\n", rElem, rElem )
	rElem.SetInt(30)
}

func main() {

	// 通过反射修改 num int的值
	// 修改 stu Studnet的值

	num := 10
	reflectTest(&num)
	fmt.Println(num)
}