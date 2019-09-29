package main
import (
	"fmt"
	"reflect"
	"strings"
)

// 获得结构体的tag标签，遍历字段的值，修改字段的值
func ToJsonStr(b interface{}) (string, error) {

	// 匿名函数，捕获异常并返回
	defer func() {
		err := recover();
		if err != nil {
			fmt.Println("错误，错误原因：", err)
		}
	}()

	rVal := reflect.ValueOf(b)
	var rType reflect.Type
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
		temp := rVal.Field(i)
		str += "\":" + fmt.Sprintf("%v", temp) + "\","
	}

	str = strings.TrimRight(str, ",") + "}"
	return str, nil
}

func ToJsonStrSlice(b interface{}) string {

	rVal := reflect.ValueOf(b)
	rKind := rVal.Kind()
	// rType := rVal.Type()

	str := "["
	if rKind == reflect.Slice {
		for i := 0; i < rVal.Len(); i++ {
			switch rVal.Index(i).Kind() {
			case reflect.Slice:
				str += ToJsonStrSlice(rVal.Index(i))
			case reflect.String:
				str += "\"" + fmt.Sprintf("%v", rVal.Index(i)) +"\","
			default:
				str += fmt.Sprintf("%v", rVal.Index(i))
			}
			temp := reflect.ValueOf(rVal.Index(i))
			if temp.Kind() == reflect.Struct {
				fmt.Println("true")
				fmt.Printf("%v\n", temp.Field(0))
			}
			fmt.Printf("%T\n", temp)
			// fmt.Println("kind = ", temp.Kind(), reflect.ValueOf(rVal.Index(i)).Kind())
		}
		str = strings.TrimRight(str, ",")
	}
	return str + "]"
}

func Test() {
	stu := Student{
		Name : "张三",
		Age : 18,
		Address : []string{
			"江西",
			"九江",
			"都昌",
		},
	}

	stuJson, err := ToJsonStr(stu)
	if err != nil {
		fmt.Println(err)
		return 
	} 
	fmt.Println(stuJson)
}

type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Address []string
}
func main() {

	s := []interface{} {
		"中学生",
		"小学生",
		"大学生",
		1,
	}

	fmt.Println(s[0])
	fmt.Println(ToJsonStrSlice(s))
}