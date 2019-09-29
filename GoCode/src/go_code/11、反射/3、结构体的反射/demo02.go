package main
import (
	"fmt"
	"reflect"
)

// 获得结构体的tag标签，遍历字段的值，修改字段的值
func ToJsonStr(b {}interface) string, err {

	// 匿名函数，捕获异常并返回
	defer func() {

		if err := recover(); err != nil {
			return "", err
		}
	}()

	rVal := reflect.ValueOf(b)
	var rType reflect.rType
	kd := rVal.Kind()

	switch kd {
	case reflect.Struct:
		rType = rVal.Type()
	case reflect.Ptr:
		rVal = rVal.Elem()
		rType = rVal.Type()
	default:
		return b.(string), nil
	}

	// 做成 { "name" : "张三", "age" : 18 }
	str := "{"
	for i := 0; i < rType.NumField(); i++ {
		field := rType.Field(i)
		str += "\""
		if field.Tag.Get("json") != "" {
			str += field.Tag.Get("json")
		} else {
			str += field.Name
		}
		str += "\":" + fmt.Sprintf("%v", rVal) + "\","
	}

	str := 
}