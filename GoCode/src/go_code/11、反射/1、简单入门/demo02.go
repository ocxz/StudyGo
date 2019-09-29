package main
import (
	"fmt"
	"reflect"
)

// 定义一个结构体
type Student struct {
	Name string
	Age int
}

func reflectTest(b interface{}) {

	rType := reflect.TypeOf(b)
	fmt.Printf("rType = %v，类型是：%T\n", rType, rType)

	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue = %v，类型是：%T\n", rValue, rValue)

	// 可以通过rType、rValue两种方式获得
	rKind := rValue.Kind()
	fmt.Printf("通过rValue的方式：rKind = %v，类型是：%T\n", rKind, rKind)
	rKind2 := rType.Kind()
	fmt.Printf("通过rType的方式：rKind2 = %v，类型是：%T\n", rKind2, rKind2)

	iv := rValue.Interface()
	fmt.Printf("iv = %v， iv的类型是：%T\n", iv, iv)

	stu, ok := iv.(Student)
	if ok {
		fmt.Println(stu.Name)
	}
}
func main() {

	// 演示反射机制对结构体的操作
	stu := Student {
		Name : "tom",
		Age : 20,
	}

	reflectTest(stu)
}