package main
import (
	"fmt"
)

func main() {

	// 三种声明方式
	// 第一种先声明、后分配内存、再使用
	var citys map[string]string
	citys = make(map[string]string, 10)
	citys["no1"] = "北京"
	fmt.Println(citys)

	// 第二种声明时分配内存、然后直接使用
	var persons = make(map[string]string, 10)
	persons["张三"] = "老师"
	fmt.Println(persons)

	// 第三种、直接声明、分配内存、初始化一起
	var fruits = map[string]string { "no1" : "苹果", "no2" : "香蕉", }
	fruits["no3"] = "梨"
	fmt.Println(fruits)

	// map的多层嵌套使用
	students := make(map[string]map[string]string)  // 第一层map
	students["no1"] = make(map[string]string, 2)    // 第二层map
	students["no1"]["name"] = "张三"
	students["no1"]["sex"] = "男"

	students["no2"] = make(map[string]string, 2)
	students["no2"]["name"] = "丽丝"
	students["no2"]["sex"] = "女"
	fmt.Println(students)





}