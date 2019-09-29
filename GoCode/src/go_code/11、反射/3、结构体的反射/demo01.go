package main
import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name string `json:"name" tag:"测试标签"`
	Age int `json:"age"`
	Score float32 `tag:"测试标签"`
	Sex string
}

func (m *Monster) Print() {
	fmt.Println("-----------start-----------")
	fmt.Println(*m)
	fmt.Println("-----------end-----------")
}

func (m *Monster) GetSum(n1 int, n2 int) int {
	return n1 + n2
}

func (m *Monster) Set(name string, age int, score float32, sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}


func main() {

	// 声明定义结构体变量值
	m := &Monster {
		Name : "黄鼠狼",
		Age : 400,
		Score : 30.8,
		Sex : "男",
	}

	// 获取三个反射中常用的值
	prVal := reflect.ValueOf(m)   // 执行结构体指针的value
	rVal := prVal.Elem()   // prVal指向的结构体
	rType := rVal.Type()
	rKind := rVal.Kind()

	// 判断是否为结构体
	if rKind != reflect.Struct {
		fmt.Println("不是一个结构体，没有必须往下走了")
		return
	}

	// 获取该结构体有多少个字段
	fieldCounts := rVal.NumField()	
	fmt.Printf("%v结构体共有%v个字段\n", rType, fieldCounts)

	// 遍历结构体中所有字段
	for i := 0; i < fieldCounts; i++ {
		fieldKind := rType.Field(i)
		fieldValue := rVal.Field(i)
		fmt.Printf("%v = %v\t值的类型是：%v\n", fieldKind.Name,
		fieldValue, fieldValue.Type())

		// 获取字段json标签的值
		if tagVal := fieldKind.Tag.Get("tag"); tagVal != "" {
			fmt.Printf("%v具有一个tag标签，tag标签的值为：%v\n", fieldKind.Name, tagVal)
		}
	}

		// 获取结构体由多少个方法
		methodCount := prVal.NumMethod()
		fmt.Printf("%v结构体具有%v个方法\n", rType, methodCount)
		
		// 调用prVal的第1个方法，按照方法名的ascii进行排序
		prVal.Method(1).Call(nil)

		pars := []int{3, 2}
		rpars := []reflect.Value{}
		for _, par := range pars {
			fmt.Println("par = ", par)
			rpars = append(rpars, reflect.ValueOf(par))
		}
		
		// 调用(m *Monster)GetSum(int, int) int方法
		md := prVal.Type().Method(0)
		fmt.Println("方法名：", md.Name)
		res := prVal.Method(0).Call(rpars)
		fmt.Printf("%v.GetSum(%v, %v) = %v", rType, pars[0], pars[1], res[0].Int())
		
}