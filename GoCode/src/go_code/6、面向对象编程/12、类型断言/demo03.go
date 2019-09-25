package main
import (
	"fmt"
)

// 最好写在使用前面
type Student struct {
	Name string
	Age int
}

// 编写一个函数，可以判断输入的参数是什么类型，使用类型断言的方式
func TypeJude(items ...interface{}) {
	
	for k, v := range items {
		switch v.(type) {
			case bool:
				fmt.Printf("第%v个参数是bool类型，值是%v\n",k, v)
			case float32:
				fmt.Printf("第%v个参数是float32类型，值是%v\n",k, v)
			case float64:
				fmt.Printf("第%v个参数是float64类型，值是%v\n",k, v)
			case int, int32, int64:
				fmt.Printf("第%v个参数是整数类型，值是%v\n",k, v)
			case string:
				fmt.Printf("第%v个参数是字符串类型，值是%v\n",k, v)
			// Student类型
			case Student:
				fmt.Printf("第%v个参数是Student类型，值是%v\n",k, v)
			// *Student类型
			case *Student:
				fmt.Printf("第%v个参数是*Student类型，值是%v\n",k, v)
			default:
				fmt.Printf("第%v个参数是类型不确定，值是%v\n",k, v)
		}
	}
}

func main() {

	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 2
	var name string = "tom"
	stu1 := Student{"张三", 18}
	stu2 := &Student{"李四", 20}
	TypeJude(n1, n2, n3, name, stu1, stu2)
}