package main
import (
	"reflect"
	"fmt"
)

// 专门演示反射的函数
func reflectTest(b interface{}) {

	// 通过反射获取传入变量的type、kin、值
	rType := reflect.TypeOf(b)
	fmt.Println("rType = ", rType)

	// 获取到reflect.Value()
	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue = %v, rValue的类型是：%T\n", rValue, rValue)

	// 获取到reflect.Value真正的值
	c := 10 + rValue.Int()
	fmt.Println(c)

	// 将rValue转成interface{} 
	iv := rValue.Interface()
	// 将interfacce{} 通过断言转成需要转成的类型
	num2 := iv.(int)
	fmt.Println("num2 = ", num2)
}
func main() {

	/*
		演示对基本数据类型、interface{}、reflect.Value()进行反射的基本操作
		通过反射拿到基本数据类型的一系列信息
	*/

	intNum := 100
	reflectTest(intNum)


}