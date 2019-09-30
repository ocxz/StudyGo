package regiorm
import (
	"reflect"
	_ "errors"
)

// 将结构体转成[]interface类型数据
func GetParams(s interface{}) (res []interface{}) {

	// s 要是结构体指针
	rValue := reflect.ValueOf(s)
	rKind := rValue.Kind()
	var rType reflect.Type

	switch rKind {
	case reflect.Ptr:
		rValue = rValue.Elem()
		fallthrough
	case reflect.Struct:
		rType = rValue.Type()
	default:
		return append(res, s)
	}

	// 获得结构体的字段名，字段名 + 字段值
	fieldCount := rValue.NumField()
	for i := 0; i < fieldCount; i ++ {

		// 获取字段名称
		name := rType.Field(i).Name
		tag := rType.Field(i).Tag.Get("orm")
		// 获取字段的值
		val := rValue.Field(i).Interface()
		if tag != "" {
			name = tag
		}
		res = append(res, name, val)
	}
	return res
}

// 将data中的数据，转化为结构体
// 两种传参，1、结构体指针，但必须
// func GetHash(ptr interface{}, data []interface{}) error {

// 	rValue := reflect.ValueOf(ptr)
// 	rKind := rValue.Kind()
// 	var rType reflect.Type

// 	// switch rKind
// 	return errors.New("错误")
// }

// 判断元素是否存在于切片中，如果不存在返回负一
func ExistSlice(val interface{}, slice []string) int {

	for k, v := range slice {
		if val == v {
			return k
		}
	}
	return -1
}
