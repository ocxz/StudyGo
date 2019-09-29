package main
import (
	"fmt"
	"reflect"
)

func main() {

	// var v float64 = 1.2
	// 通过反射获得value、Type、Kind
	// 将value转成interface{}
	// 通过interface{}，转成float64

	var v float64 = 1.2
	rValue := reflect.ValueOf(v)
	fmt.Println("rValue = ", rValue)   // 

	rType := rValue.Type()
	fmt.Println("rType = ", rType)

	rKind := rValue.Kind()
	fmt.Println("rKind = ", rKind)

	// 编译时是分析不了iv的类型，因为它是interface{}
	// 但是通过fmt.Printf(iv)打印时，底层是可分析出iv的实际类型
	// 因此 iv := rValue.Interface() + 2.3  错误的，编译通不过
	// mismatched types interface {} and float64) 
	iv := rValue.Interface()
	fmt.Printf("iv = %v，iv的type = %T\n", iv, iv)

	v2 := iv.(float64) + 3.6
	fmt.Printf("v2 = %v，v2的type = %T\n", v2, v2)

	
}